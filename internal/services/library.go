package services

import (
	"context"

	"fanchann/library/internal/models/web"
)

type ILibraryServices interface {
	AddNewBook(ctx context.Context, formData web.AddNewBooks) web.AddNewBooks
	UpdateBook(ctx context.Context, formData web.UpdateBook) (web.UpdateBook, error)
	DeleteBook(ctx context.Context, bookID int) error
	FindByIdBook(ctx context.Context, bookID int) (web.BooksResponse, error)
	FindByIdAuthor(ctx context.Context, authorID int) (web.AuthorsResponse, error)
	FindAllBook(ctx context.Context) []web.BooksResponse
}
