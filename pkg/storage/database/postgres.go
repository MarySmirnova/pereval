package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/MarySmirnova/pereval/internal/config"
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

type TXpg struct {
	tx *pgx.Tx
}

func (s *Storage) NewTXpg() (*TXpg, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &TXpg{
		tx: &tx,
	}, nil
}
