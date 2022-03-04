package models

type Images struct {
	Img []*Image `json:"images"`
}

type Image struct {
	URL   []string `json:"url"`
	Title string   `json:"title"`
}

type ImgsAdded map[string][]int
