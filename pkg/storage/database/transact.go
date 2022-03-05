package database

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/MarySmirnova/pereval/internal/data"
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

func (t *TXpg) putImages(data *models.Pereval, images *models.Images) (*models.ImgsAdded, error) {
	imgsAdded := make(models.ImgsAdded)
	dateAdded := data.AddTime

	for key, imgs := range images.Img {
		imgsAdded[key] = []int{}

		for _, img := range imgs {
			if img.URL == "" {
				continue
			}
			file, err := http.Get(img.URL)
			if err != nil {
				return nil, err
			}
			body, err := io.ReadAll(file.Body)
			if err != nil {
				return nil, err
			}

			qu := fmt.Sprintf(`INSERT INTO public.pereval_images (date_added, img, title)
			VALUES ('%s', '%v'::bytea, '%s')
			RETURNING id;`, dateAdded, body, img.Title)

			var id int
			err = t.tx.QueryRow(ctx, qu).Scan(&id)
			if err != nil {
				return nil, err
			}
			imgsAdded[key] = append(imgsAdded[key], id)
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

func (t *TXpg) getData(id int) ([]byte, error) {
	query := fmt.Sprintf(`SELECT raw_data
	FROM public.pereval_added
	WHERE id = %d;`, id)

	var data []byte
	row := t.tx.QueryRow(ctx, query)
	err := row.Scan(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t *TXpg) getImagesID(id int) (*models.ImgsAdded, error) {
	query := fmt.Sprintf(`SELECT images
	FROM public.pereval_added
	WHERE id = %d;`, id)

	imgMap := make(models.ImgsAdded)
	var imgJson []byte
	row := t.tx.QueryRow(ctx, query)
	err := row.Scan(&imgJson)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(imgJson, &imgMap)
	if err != nil {
		return nil, err
	}

	return &imgMap, nil
}

func (t *TXpg) convertImages(imgMap *models.ImgsAdded, pereval *data.Pereval) error {
	query := `SELECT img, title
	FROM public.pereval_images
	WHERE id = $1`

	for key, imgID := range *imgMap {
		if _, ok := pereval.Img[key]; !ok {
			pereval.Img[key] = []*data.Image{}
		}
		for _, id := range imgID {
			image := data.Image{}
			pereval.Img[key] = append(pereval.Img[key], &image)

			row := t.tx.QueryRow(ctx, query, id)
			err := row.Scan(&image.Img, &image.Title)
			if err != nil {
				return err
			}
			image.ID = id
		}
	}

	return nil
}

func (t *TXpg) delImages(oldImgs *models.ImgsAdded) error {
	query := `DELETE FROM public.pereval_images
	WHERE id = $1`

	for _, imgID := range *oldImgs {
		for _, id := range imgID {
			_, err := t.tx.Exec(ctx, query, id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *TXpg) updateData(data *models.Pereval, imgs *models.ImgsAdded, id int) error {
	jsonImg, err := json.Marshal(imgs)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	query := `UPDATE public.pereval_added
	SET date_added = $1, raw_data = $2, images = $3
	WHERE id = $4;`

	_, err = t.tx.Exec(ctx, query, data.AddTime, jsonData, jsonImg, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *TXpg) getPereval(id int) (*data.Pereval, error) {
	dat, err := t.getData(id)
	if err != nil {
		return nil, err
	}

	imgsID, err := t.getImagesID(id)
	if err != nil {
		return nil, err
	}

	pereval := data.Pereval{
		Img: make(map[string][]*data.Image),
	}

	err = json.Unmarshal(dat, &pereval)
	if err != nil {
		return nil, err
	}

	err = t.convertImages(imgsID, &pereval)
	if err != nil {
		return nil, err
	}

	return &pereval, nil
}

func (t *TXpg) selectDataIDs(user map[string]string) ([]int, error) {
	ids := []int{}

	for key, param := range user {
		if len(ids) == 0 {
			qu := fmt.Sprintln(`SELECT id FROM public.pereval_added 
			WHERE raw_data::TEXT SIMILAR TO '%"` + key + `":"` + param + `"%';`)
			rows, err := t.tx.Query(ctx, qu)
			if err != nil {
				return nil, err
			}

			for rows.Next() {
				var id int
				err = rows.Scan(&id)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}
			continue
		}

		newIds := []int{}
		for _, id := range ids {
			var scanID int

			qu := fmt.Sprintln(`SELECT id FROM public.pereval_added
			WHERE id = ` + strconv.Itoa(id) +
				`AND raw_data::TEXT SIMILAR TO '%"` + key + `":"` + param + `"%';`)
			row := t.tx.QueryRow(ctx, qu)
			err := row.Scan(&scanID)
			if err != nil {
				if err == pgx.ErrNoRows {
					continue
				}
				return nil, err
			}
			newIds = append(newIds, scanID)
		}

		ids = newIds
	}

	return ids, nil
}
