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
		Book_title  string
		Inserted_At string
		Updated_At  string
	}

	Books_Information struct {
		Book_id   int
		Author_id int
	}
)
