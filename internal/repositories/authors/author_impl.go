package authors

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"fanchann/library/internal/models/domain"
	"fanchann/library/pkg/utils"
)

type AuthorRepoImpl struct{}

func NewAuthorRepoImpl() IAuthorsRepositories {
	return &AuthorRepoImpl{}
}

func (author *AuthorRepoImpl) Add(ctx context.Context, tx *sql.Tx, authorData *domain.Author) domain.Author {
	addQuery := `insert into authors(author_name) values (?)`
	id, err := tx.ExecContext(ctx, addQuery, authorData.Author_Name)
	utils.LogErrorWithPanic(err)
	intId, _ := id.LastInsertId()
	authorData.Author_Id = int(intId)
	return *authorData
}

func (author *AuthorRepoImpl) Update(ctx context.Context, tx *sql.Tx, authorData *domain.Author) domain.Author {
	updateQuery := `update authors set author_name = ? where author_id = ?`
	id, err := tx.ExecContext(ctx, updateQuery, authorData.Author_Name, authorData.Author_Id)
	utils.LogErrorWithPanic(err)
	intId, _ := id.LastInsertId()
	authorData.Author_Id = int(intId)
	return *authorData
}

func (author *AuthorRepoImpl) FindAuthorById(ctx context.Context, tx *sql.Tx, authorId int) (domain.Author, error) {
	findByIdQuery := `select author_id,author_name from authors where author_id = ?`
	rows, err := tx.QueryContext(ctx, findByIdQuery, authorId)
	utils.LogErrorWithPanic(err)
	defer rows.Close()

	if rows.Next() {
		author := domain.Author{}
		err := rows.Scan(&author.Author_Id, &author.Author_Name)
		utils.LogErrorWithPanic(err)
		return author, nil
	}
	notFoundMsg := fmt.Sprintf("author with id %d not found", authorId)
	return domain.Author{}, errors.New(notFoundMsg)
}

func (author *AuthorRepoImpl) FindAllAuthor(ctx context.Context, tx *sql.Tx) []domain.Author {
	authors := []domain.Author{}
	findAllQuery := `select author_id,author_name from authors`
	rows, err := tx.QueryContext(ctx, findAllQuery)
	utils.LogErrorWithPanic(err)
	defer rows.Close()

	for rows.Next() {
		author := domain.Author{}
		err := rows.Scan(&author.Author_Id, &author.Author_Name)
		utils.LogErrorWithPanic(err)
		authors = append(authors, author)
	}
	return authors
}

func (author *AuthorRepoImpl) FindAuthorByName(ctx context.Context, tx *sql.Tx, name string) domain.Author {
	findNameQuery := `select author_id,author_name from authors where author_name = ?`
	row, err := tx.QueryContext(ctx, findNameQuery, name)
	utils.LogErrorWithPanic(err)
	defer row.Close()

	data := domain.Author{}
	if row.Next() {
		errScan := row.Scan(&data.Author_Id, &data.Author_Name)
		utils.LogErrorWithPanic(errScan)
		return data
	}
	return domain.Author{}
}
