package database

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/MarySmirnova/pereval/internal/config"
	"github.com/MarySmirnova/pereval/internal/data"
	"github.com/MarySmirnova/pereval/pkg/storage/models"
)

var ctx context.Context = context.Background()

type Storage struct {
	db *pgxpool.Pool
}

func NewDBpg(cfg config.Postgres) (*Storage, error) {
	connect := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := pgxpool.Connect(ctx, connect)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) GetPGXpool() *pgxpool.Pool {
	return s.db
}

func (s *Storage) SubmitData(data *data.Pereval) (id int, err error) {
	var rawData *models.Pereval
	var images []*models.Image

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	dateAdded := data.AddTime

	for _, img := range data.Images {
		file, err := http.Get(img.URL)
		if err != nil {
			return 0, err
		}
		body, err := ioutil.ReadAll(file.Body)
		if err != nil {
			return 0, err
		}

		qu := fmt.Sprintf(`INSERT INTO public.pereval_images (date_added, img)
		VALUES ('%s', '%v'::bytea)
		RETURNING id;`, dateAdded, body)

		image := &models.Image{}
		err = tx.QueryRow(ctx, qu).Scan(&image.IDimg)
		if err != nil {
			return 0, err
		}
		image.Title = img.Title
		image.URL = img.URL
		images = append(images, image)
	}

	conv, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	err = json.NewDecoder(bytes.NewBuffer(conv)).Decode(&rawData)
	if err != nil {
		return 0, err
	}

	jsonImg, err := json.Marshal(images)
	if err != nil {
		return 0, err
	}
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO public.pereval_added (date_added, raw_data, images, status)
	VALUES ($1, $2, $3, 'new')
	RETURNING id;`

	err = tx.QueryRow(ctx, query, dateAdded, jsonData, jsonImg).Scan(&id)
	if err != nil {
		return 0, err
	}

	tx.Commit(ctx)
	return id, nil
}
