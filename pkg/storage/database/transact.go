package database

import (
	"fmt"
	"io"
	"net/http"

	"github.com/MarySmirnova/pereval/internal/data"
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

func (t *TXpg) putImages(data *data.Pereval, imgs *data.Images) error {
	dateAdded := data.AddTime

	for _, img := range imgs.Img {
		file, err := http.Get(img.URL)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(file.Body)
		if err != nil {
			return err
		}

		qu := fmt.Sprintf(`INSERT INTO public.pereval_images (date_added, img)
		VALUES ('%s', '%v'::bytea)
		RETURNING id;`, dateAdded, body)

		err = t.tx.QueryRow(ctx, qu).Scan(&img.IDimg)
		if err != nil {
			return err
		}
	}

	return nil
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
