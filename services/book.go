package services

import (
	"errors"
	"go-crud-backend/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

var books = []models.Book{
	{ID: "1", Title: "Rich Dad Poor Dad", Author: "Robert Kiyosaki", Price: 12.99},
}

func GetAllBooks() []models.Book {
	return books
}

func GetBookByID(id string) (models.Book, error) {
	for _, b := range books {
		if b.ID == id {
			return b, nil
		}
	}
	return models.Book{}, errors.New("book not found")
}

func CreateBook(newBook models.Book) (models.Book, error) {
	if err := validate.Struct(newBook); err != nil {
		return models.Book{}, err
	}

	for _, b := range books {
		if b.ID == newBook.ID {
			return models.Book{}, errors.New("a book with this ID already exists")
		}
	}

	books = append(books, newBook)
	return newBook, nil
}

func UpdateBook(id string, updatedBook models.UpdateBookInput) (models.Book, error) {
	for i, b := range books {
		if b.ID == id {
			if updatedBook.Title != nil {
				books[i].Title = *updatedBook.Title
			}
			if updatedBook.Author != nil {
				books[i].Author = *updatedBook.Author
			}
			if updatedBook.Price != nil {
				books[i].Price = *updatedBook.Price
			}

			return books[i], nil
		}
	}
	return models.Book{}, errors.New("book not found")
}

func DeleteBook(id string) error {
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
