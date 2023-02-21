package bookingbook

import (
	"fmt"
)

type BooksController interface {
	GetBooks(query map[string]string) BooksResponse
	SubmitBooks(body BodySubmitBook) SubmitResponse
}

type booksController struct {
	booksUseCase BooksUseCase
}

func (c booksController) GetBooks(query map[string]string) BooksResponse {
	res := BooksResponse{}
	result, err := c.booksUseCase.GetBooks(query)
	if err != nil {
		fmt.Printf("error")
		return BooksResponse{}
	}
	res.Success = true
	res.MessageTitle = "Get list book by subject"
	res.Message = "Success get list book"
	res.Data = result

	return res
}

func (c booksController) SubmitBooks(query BodySubmitBook) SubmitResponse {
	res := SubmitResponse{}
	result, err := c.booksUseCase.SubmitBooks(query)
	if err != nil {
		fmt.Printf("error")
		return SubmitResponse{}
	}
	res.Success = result
	res.MessageTitle = "Request pick book"
	res.Message = "Success Submit"
	res.Data = query

	return res
}
