package authors

import (
	"context"
	"database/sql"

	"fanchann/library/internal/models/domain"
)

type IAuthorsRepositories interface {
	Add(ctx context.Context, tx *sql.Tx, authorData *domain.Author) domain.Author
	Update(ctx context.Context, tx *sql.Tx, authorData *domain.Author) domain.Author
	FindAuthorById(ctx context.Context, tx *sql.Tx, authorId int) (domain.Author, error)
	FindAllAuthor(ctx context.Context, tx *sql.Tx) []domain.Author
	FindAuthorByName(ctx context.Context, tx *sql.Tx, name string) error
}
