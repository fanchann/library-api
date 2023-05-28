package booksinformation

import (
	"context"
	"database/sql"

	"fanchann/library/internal/models/domain"
)

type IBooksInformation interface {
	Insert(ctx context.Context, tx *sql.Tx, dataBooks *domain.Books_Information) domain.Books_Information
	Delete(ctx context.Context, tx *sql.Tx, bookID int) error
	FindBookIdWithAuthor(ctx context.Context, tx *sql.Tx, bookId int) (domain.BookWithAuthor, error)
	FindAuthorIdWithBooks(ctx context.Context, tx *sql.Tx, authorId int) []domain.BookLists
	FindAllBooksAndAuthor(ctx context.Context, tx *sql.Tx) []domain.BookWithAuthor
}
