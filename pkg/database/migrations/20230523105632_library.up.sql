CREATE TABLE
    books (
        book_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (book_id),
        book_title VARCHAR(100) NOT NULL,
        inserted_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP
    );

CREATE TABLE
    authors (
        author_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (author_id),
        author_name VARCHAR(100) NOT NULL,
        inserted_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP
    );

CREATE TABLE
    books_information (
        book_id INT UNSIGNED,
        author_id INT UNSIGNED,
        CONSTRAINT fk_booksInfo_books FOREIGN KEY (book_id) REFERENCES books (book_id),
        CONSTRAINT fk_booksInfo_authors FOREIGN KEY (author_id) REFERENCES authors (author_id)
    );