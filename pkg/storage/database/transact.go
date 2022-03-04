package database

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MarySmirnova/pereval/pkg/storage/models"
	"github.com/jackc/pgx/v4"
)

type TXpg struct {
	tx pgx.Tx
}

func (s *Storage) NewTXpg() (*TXpg, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &TXpg{
		tx: tx,
	}, nil
}

func (t *TXpg) putImages(data *models.Pereval, imgs *models.Images) (*models.ImgsAdded, error) {
	imgsAdded := make(models.ImgsAdded)
	dateAdded := data.AddTime

	for _, img := range imgs.Img {
		imgsAdded[img.Title] = []int{}

		for _, url := range img.URL {
			file, err := http.Get(url)
			if err != nil {
				return nil, err
			}
			body, err := io.ReadAll(file.Body)
			if err != nil {
				return nil, err
			}

			qu := fmt.Sprintf(`INSERT INTO public.pereval_images (date_added, img)
			VALUES ('%s', '%v'::bytea)
			RETURNING id;`, dateAdded, body)

			var id int
			err = t.tx.QueryRow(ctx, qu).Scan(&id)
			if err != nil {
				return nil, err
			}
			imgsAdded[img.Title] = append(imgsAdded[img.Title], id)
		}
	}

	return &imgsAdded, nil
}

func (t *TXpg) putData(data *models.Pereval, imgs *models.ImgsAdded) (int, error) {
	jsonImg, err := json.Marshal(imgs)
	if err != nil {
		return 0, err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO public.pereval_added (date_added, raw_data, images, status)
	VALUES ($1, $2, $3, $4)
	RETURNING id;`

	var id int
	err = t.tx.QueryRow(ctx, query, data.AddTime, jsonData, jsonImg, statusNew).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TXpg) getStatus(id int) (string, error) {
	query := fmt.Sprintf(`SELECT status 
	FROM public.pereval_added
	WHERE id = %d;`, id)

	var status string
	row := t.tx.QueryRow(ctx, query)
	err := row.Scan(&status)
	if err != nil {
		return "", err
	}
	return status, nil
}
