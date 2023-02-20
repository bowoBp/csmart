package bookingbook

import (
	"encoding/json"
	"net/http"
)

type BooksUseCase interface {
	GetBooks(query map[string]string) ([]Book, error)
}

type bookUseCase struct {
}

func (uc bookUseCase) GetBooks(query map[string]string) ([]Book, error) {
	return getListBooksBySubject(query)
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
