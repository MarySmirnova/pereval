package models

import "github.com/MarySmirnova/pereval/internal/data"

type Pereval struct {
	ID          string      `json:"id"`
	BeautyTitle string      `json:"beautyTitle"`
	Title       string      `json:"title"`
	OtherTitles string      `json:"other_titles"`
	Connect     string      `json:"connect"`
	AddTime     string      `json:"add_time"`
	User        data.User   `json:"user"`
	Coords      data.Coords `json:"coords"`
	Type        string      `json:"type"`
	Level       data.Level  `json:"level"`
}
