package data

import (
	"fmt"
	"net/mail"
	"strconv"
)

func Validate(data *Pereval) error {
	if data.Title == "" {
		return fmt.Errorf("pass must be specified")
	}

	if !validCoords(data.Coords.Height, data.Coords.Latitude, data.Coords.Longitude) {
		return fmt.Errorf("coordinates must be specified")
	}

	if !validEmail(data.User.Email) {
		return fmt.Errorf("invalid email entered")
	}

	if !vlaidImgs(data) {
		return fmt.Errorf("please add a photo")
	}

	return nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validCoords(coords ...string) bool {
	for _, coord := range coords {
		if _, err := strconv.ParseFloat(coord, 64); err != nil {
			return false
		}
	}
	return true
}

func vlaidImgs(data *Pereval) bool {
	for _, imgs := range data.Img {
		for _, img := range imgs {
			if img.URL != "" {
				return true
			}
		}
	}

	return false
}
