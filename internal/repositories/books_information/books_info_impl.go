package booksinformation

import (
	"context"
	"database/sql"

	"fanchann/library/internal/models/domain"
	"fanchann/library/pkg/utils"
)

type BooksInfoRepoImpl struct{}

func NewBooksInfoRepoImpl() IBooksInformation {
	return &BooksInfoRepoImpl{}
}

func (bookInfo *BooksInfoRepoImpl) Insert(ctx context.Context, tx *sql.Tx, dataBooks *domain.Books_Information) domain.Books_Information {
	addQuery := `insert into books_information(book_id,author_id) values(?,?)`
	_, err := tx.ExecContext(ctx, addQuery, dataBooks.Book_id, dataBooks.Author_id)
	utils.LogErrorWithPanic(err)
	return *dataBooks
}

func (bookInfo *BooksInfoRepoImpl) Delete(ctx context.Context, tx *sql.Tx, dataBooks *domain.Books_Information) error {
	deleteQuery := `delete from books_information where book_id = ? AND author_id = ?`
	_, err := tx.ExecContext(ctx, deleteQuery, dataBooks.Book_id, dataBooks.Author_id)
	if err != nil {
		return err
	}
	return nil
}
