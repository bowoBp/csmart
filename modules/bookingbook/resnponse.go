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

type BodySubmitBook struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	PickupDate string `json:"pickup_date"`
	TotalBook  int    `json:"total_book"`
	Books      []Book `json:"books"`
}

type SubmitResponse struct {
	dto.ResponseMeta
	Data BodySubmitBook `json:"data"`
}
