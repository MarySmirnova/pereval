package models

type Pereval struct {
	ID          string  `json:"id"`
	BeautyTitle string  `json:"beautyTitle"`
	Title       string  `json:"title"`
	OtherTitles string  `json:"other_titles"`
	Connect     string  `json:"connect"`
	AddTime     string  `json:"add_time"`
	User        User    `json:"user"`
	Coords      Coords  `json:"coords"`
	Type        string  `json:"type"`
	Level       Level   `json:"level"`
	Images      []Image `json:"images"`
}

type User struct{}

type Coords struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Height    string `json:"height"`
}

type Level struct {
	Winter string `json:"winter"`
	Summer string `json:"summer"`
	Autumn string `json:"autumn"`
	Spring string `json:"spring"`
}

type Image struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}
