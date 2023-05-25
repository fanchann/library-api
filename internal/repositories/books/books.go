package books

import (
	"context"
	"database/sql"

	"fanchann/library/internal/models/domain"
)

type IBooksRepositories interface {
	Add(ctx context.Context, tx *sql.Tx, bookData *domain.Books) domain.Books
	Update(ctx context.Context, tx *sql.Tx, bookData *domain.Books) domain.Books
	FindById(ctx context.Context, tx *sql.Tx, bookID int) (domain.Books, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Books
	Delete(ctx context.Context, tx *sql.Tx, bookID int)
}
