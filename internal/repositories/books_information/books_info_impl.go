package booksinformation

import (
	"context"
	"database/sql"
	"errors"

	"fanchann/library/internal/models/domain"
	"fanchann/library/pkg/utils"
)

type BooksInfoRepoImpl struct{}

func NewBooksInfoRepoImpl() IBooksInformation {
	return &BooksInfoRepoImpl{}
}

func (bookInfo *BooksInfoRepoImpl) Insert(ctx context.Context, tx *sql.Tx, dataBooks *domain.Books_Information) domain.Books_Information {
	addQuery := `insert into books_information(book_id,author_id) values (?,?)`
	_, err := tx.ExecContext(ctx, addQuery, dataBooks.Book_id, dataBooks.Author_id)
	utils.LogErrorWithPanic(err)
	return *dataBooks
}

func (bookInfo *BooksInfoRepoImpl) Delete(ctx context.Context, tx *sql.Tx, bookID int) error {
	deleteQuery := `delete from books_information where book_id = ?`
	_, err := tx.ExecContext(ctx, deleteQuery, bookID)
	if err != nil {
		return err
	}
	return nil
}

func (bookInfo *BooksInfoRepoImpl) FindBookIdWithAuthor(ctx context.Context, tx *sql.Tx, bookId int) (domain.BookWithAuthor, error) {
	findQuery := `select a.author_name, b.book_title
	from authors as a
			 join books_information bi on a.author_id = bi.author_id
			 join books b on bi.book_id = b.book_id
	where b.book_id = ?`
	rows, err := tx.QueryContext(ctx, findQuery, bookId)
	utils.LogErrorWithPanic(err)
	defer rows.Close()

	if rows.Next() {
		dataBook := domain.BookWithAuthor{}
		err := rows.Scan(&dataBook.Author_Name, &dataBook.Book_Title)
		utils.LogErrorWithPanic(err)
		return dataBook, nil
	}
	return domain.BookWithAuthor{}, errors.New("data not found")
}

func (bookInfo *BooksInfoRepoImpl) FindAuthorIdWithBooks(ctx context.Context, tx *sql.Tx, authorId int) []domain.BookLists {
	findQuery := `select b.book_title
	from authors as a
			 join books_information bi on a.author_id = bi.author_id
			 join books b on bi.book_id = b.book_id
	where a.author_id = ?`
	rows, err := tx.QueryContext(ctx, findQuery, authorId)
	utils.LogErrorWithPanic(err)
	defer rows.Close()

	bookLists := []domain.BookLists{}
	book := domain.BookLists{}

	for rows.Next() {
		err := rows.Scan(&book.Book_Title)
		utils.LogErrorWithPanic(err)
		bookLists = append(bookLists, book)
	}
	return bookLists
}

func (bookInfo *BooksInfoRepoImpl) FindAllBooksAndAuthor(ctx context.Context, tx *sql.Tx) []domain.BookWithAuthor {
	var booksAndAuthorDomain []domain.BookWithAuthor
	getAllQuery := `select b.book_title, a.author_name
	from books as b
			 join books_information bi on b.book_id = bi.book_id
			 join authors a on a.author_id = bi.author_id`
	rows, err := tx.QueryContext(ctx, getAllQuery)
	utils.LogErrorWithPanic(err)
	defer rows.Close()

	for rows.Next() {
		var bookAndAuthorDomain domain.BookWithAuthor
		errScan := rows.Scan(&bookAndAuthorDomain.Book_Title, &bookAndAuthorDomain.Author_Name)
		utils.LogErrorWithPanic(errScan)
		booksAndAuthorDomain = append(booksAndAuthorDomain, bookAndAuthorDomain)
	}
	return booksAndAuthorDomain
}
