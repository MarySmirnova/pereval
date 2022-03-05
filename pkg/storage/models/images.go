package models

type Images struct {
	Img map[string][]*Image `json:"images"`
}

type Image struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type ImgsAdded map[string][]int
