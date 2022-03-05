package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/MarySmirnova/pereval/internal/config"
	"github.com/MarySmirnova/pereval/internal/data"
	"github.com/MarySmirnova/pereval/pkg/storage/models"
)

var ctx context.Context = context.Background()

const statusNew string = "new"

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

func (s *Storage) PutDataToDB(data []byte) (int, error) {
	var pereval models.Pereval
	var imgs models.Images

	err := json.Unmarshal(data, &pereval)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(data, &imgs)
	if err != nil {
		return 0, err
	}

	t, err := s.NewTXpg()
	if err != nil {
		return 0, err
	}
	defer t.tx.Rollback(ctx)

	imgAdded, err := t.putImages(&pereval, &imgs)
	if err != nil {
		return 0, err
	}

	id, err := t.putData(&pereval, imgAdded)
	if err != nil {
		return 0, err
	}

	t.tx.Commit(ctx)
	return id, nil
}

func (s *Storage) GetStatusFromDB(id int) (string, error) {
	t, err := s.NewTXpg()
	if err != nil {
		return "", err
	}
	defer t.tx.Rollback(ctx)

	status, err := t.getStatus(id)
	t.tx.Commit(ctx)
	return status, err
}

func (s *Storage) UpdateDataToDB(id int, data []byte) error {
	var pereval models.Pereval
	var imgs models.Images

	err := json.Unmarshal(data, &pereval)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &imgs)
	if err != nil {
		return err
	}

	t, err := s.NewTXpg()
	if err != nil {
		return err
	}
	defer t.tx.Rollback(ctx)

	status, err := t.getStatus(id)
	if err != nil {
		return err
	}

	if status != statusNew {
		return fmt.Errorf("wrong status: %s", status)
	}

	oldData, err := s.GetDataFromDB(id)
	if err != nil {
		return err
	}

	if oldData.User.Email != pereval.User.Email || oldData.User.Phone != pereval.User.Phone || oldData.User.Name != pereval.User.Name || oldData.User.Fam != pereval.User.Fam || oldData.User.Otc != pereval.User.Otc {
		return fmt.Errorf("full name, phone and email cannot be edited")
	}

	oldImgsID, err := t.getImagesID(id)
	if err != nil {
		return err
	}

	err = t.delImages(oldImgsID)
	if err != nil {
		return err
	}

	newImgsID, err := t.putImages(&pereval, &imgs)
	if err != nil {
		return err
	}

	err = t.updateData(&pereval, newImgsID, id)
	if err != nil {
		return err
	}

	t.tx.Commit(ctx)
	return nil
}

func (s *Storage) GetDataFromDB(id int) (*data.Pereval, error) {
	t, err := s.NewTXpg()
	if err != nil {
		return nil, err
	}
	defer t.tx.Rollback(ctx)

	pereval, err := t.getPereval(id)
	if err != nil {
		return nil, err
	}

	t.tx.Commit(ctx)
	return pereval, nil
}

func (s *Storage) GetAllDataFromDB(userParams map[string]string) ([]*data.Pereval, error) {
	t, err := s.NewTXpg()
	if err != nil {
		return nil, err
	}
	defer t.tx.Rollback(ctx)

	ids, err := t.selectDataIDs(userParams)
	if err != nil {
		return nil, err
	}

	selectPereval := []*data.Pereval{}
	for _, id := range ids {
		pereval, err := t.getPereval(id)
		if err != nil {
			return nil, err
		}
		selectPereval = append(selectPereval, pereval)
	}

	t.tx.Commit(ctx)
	return selectPereval, nil
}
