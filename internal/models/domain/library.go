package domain

type (
	Author struct {
		Author_Id   int
		Author_Name string
		Inserted_At string
		Updated_At  string
	}

	Books struct {
		Book_id     int
		Book_Title  string
		Inserted_At string
		Updated_At  string
	}

	Books_Information struct {
		Book_id   int
		Author_id int
	}

	BookLists struct {
		Book_Title string
	}

	BookWithAuthor struct {
		Author_Name string
		Book_Title  string
	}
)
