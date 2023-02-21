package bookingbook

import (
	"encoding/json"
	"net/http"
)

type BooksUseCase interface {
	GetBooks(query map[string]string) ([]Book, error)
	SubmitBooks(query BodySubmitBook) (bool, error)
}

type bookUseCase struct {
}

func (uc bookUseCase) GetBooks(query map[string]string) ([]Book, error) {
	return getListBooksBySubject(query)
}

func (uc bookUseCase) SubmitBooks(query BodySubmitBook) (bool, error) {
	datas := make([]BodySubmitBook, 0)

	datas = append(datas, BodySubmitBook{
		Name:       query.Name,
		Email:      query.Email,
		Phone:      query.Phone,
		PickupDate: query.PickupDate,
		TotalBook:  query.TotalBook,
		Books:      query.Books,
	})
	return true, nil
}

func getListBooksBySubject(query map[string]string) ([]Book, error) {
	var resLibrary *ResponseBookLibrary

	client := &http.Client{}
	param := "?limit=" + query["limit"] + "&offset=" + query["offset"]
	req, err := http.NewRequest("GET", "https://openlibrary.org/subjects/"+query["subject"]+".json"+param, nil)
	if err != nil {
		return []Book{}, err
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&resLibrary)
	if err != nil {
		return []Book{}, nil
	}

	books := make([]Book, 0)
	authors := make([]Author, 0)
	for _, item := range resLibrary.Works {
		for _, itemAuthor := range item.Authors {
			authors = append(authors, Author{
				Key:  itemAuthor.Key,
				Name: itemAuthor.Name,
			})
		}

		books = append(books, Book{
			Title:         item.Title,
			EditionNumber: item.EditionCount,
			Authors:       authors,
		})

	}

	return books, nil
}
