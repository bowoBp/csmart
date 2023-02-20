package bookingbook

import "github.com/bowoBp/csmart/dto"

type Book struct {
	Title         string   `json:"title"`
	EditionNumber int      `json:"editionNumber"`
	Authors       []Author `json:"authors"`
}

type Author struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type BooksResponse struct {
	dto.ResponseMeta
	Data []Book `json:"data"`
}
