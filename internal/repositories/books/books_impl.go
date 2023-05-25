package books

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"fanchann/library/internal/models/domain"
	"fanchann/library/pkg/utils"
)

type BooksRepoImpl struct{}

func NewBooksRepoImpl() IBooksRepositories {
	return &BooksRepoImpl{}
}

func (book *BooksRepoImpl) Add(ctx context.Context, tx *sql.Tx, bookData *domain.Books) domain.Books {
	insertQuery := `insert into books(book_title) values (?)`
	id, err := tx.ExecContext(ctx, insertQuery, bookData.Book_title)
	utils.LogErrorWithPanic(err)
	intId, _ := id.LastInsertId()
	bookData.Book_id = int(intId)
	return *bookData
}

func (book *BooksRepoImpl) Update(ctx context.Context, tx *sql.Tx, bookData *domain.Books) domain.Books {
	updateQuery := `update books set book_title = ? where id = ?`
	id, err := tx.ExecContext(ctx, updateQuery, bookData.Book_title, bookData.Book_id)
	utils.LogErrorWithPanic(err)
	intId, _ := id.LastInsertId()
	bookData.Book_id = int(intId)
	return *bookData
}

func (book *BooksRepoImpl) FindById(ctx context.Context, tx *sql.Tx, bookID int) (domain.Books, error) {
	findQuery := `select book_id,book_title from books where book_id = ?`
	row, err := tx.QueryContext(ctx, findQuery, bookID)
	utils.LogErrorWithPanic(err)

	defer row.Close()

	domainBook := domain.Books{}
	if row.Next() {
		row.Scan(&domainBook.Book_id, &domainBook.Book_title)
		return domainBook, nil

	}
	errMsg := fmt.Sprintf("book with id %d not found", bookID)
	return domain.Books{}, errors.New(errMsg)
}

func (book *BooksRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Books {
	books := []domain.Books{}

	findAllQuery := `select book_id,book_title from books`
	rows, err := tx.QueryContext(ctx, findAllQuery)
	utils.LogErrorWithPanic(err)
	defer rows.Close()
	for rows.Next() {
		book := domain.Books{}
		rows.Scan(&book.Book_id, &book.Book_title)
		books = append(books, book)
	}
	return books

}

func (book *BooksRepoImpl) Delete(ctx context.Context, tx *sql.Tx, bookID int) {
	deleteQuery := `delete from books where books = ?`
	_, err := tx.ExecContext(ctx, deleteQuery, bookID)
	utils.LogErrorWithPanic(err)
}
