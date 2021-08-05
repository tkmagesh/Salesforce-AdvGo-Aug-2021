package repository

import (
	"graphql-go-demo/graph/model"
	db "graphql-go-demo/internal/pkg/db/mysql"
	"log"
)

//CreateAuthor create's author
func CreateAuthor(author model.Author) (int64, error) {

	stmt, err := db.Db.Prepare("INSERT INTO Authors(FirstName,LastName) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	res, err := stmt.Exec(author.FirstName, author.LastName)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	defer stmt.Close()
	log.Println("Row inserted!!")
	return id, nil
}

//GetAuthorByID return author with respective id
func GetAuthorByID(id *string) (*model.Author, error) {

	stmt, err := db.Db.Prepare("select * from Authors where id=?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()
	var author model.Author
	for rows.Next() {
		err = rows.Scan(&author.ID, &author.FirstName, &author.LastName)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer rows.Close()

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &author, nil

}

//GetAllAuthors returns all authors
func GetAllAuthors() ([]*model.Author, error) {
	stmt, err := db.Db.Prepare("select * from Authors")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var authors []*model.Author
	for rows.Next() {
		var author model.Author
		rows.Scan(&author.ID, &author.FirstName, &author.LastName)
		authors = append(authors, &author)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	defer rows.Close()

	return authors, err
}

//CreateBook creates new book
func CreateBook(book model.Book) (int64, error) {
	stmt, err := db.Db.Prepare("insert into Books(Title,AuthorID) VALUES(?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(book.Title, book.Author.ID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

//GetBooksByID returns books by respective id
func GetBooksByID(id *string) (*model.Book, error) {
	stmt, err := db.Db.Prepare("select Books.ID,Books.Title,Authors.ID,Authors.FirstName,Authors.LastName from Books inner join Authors where Books.AuthorID = Authors.ID and Books.ID = ? ;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	var bookID, title, authorID, firstName, lastName string
	if rows.Next() {
		err := rows.Scan(&bookID, &title, &authorID, &firstName, &lastName)
		if err != nil {
			return nil, err
		}
	}

	book := &model.Book{
		ID:    bookID,
		Title: title,
		Author: &model.Author{
			ID:        authorID,
			FirstName: firstName,
			LastName:  lastName,
		},
	}
	defer rows.Close()
	defer stmt.Close()
	return book, nil
}

//GetAllBooks returns all Books Data
func GetAllBooks() ([]*model.Book, error) {
	var books []*model.Book
	stmt, err := db.Db.Prepare("select Books.ID,Books.Title,Authors.ID,Authors.FirstName,Authors.LastName from Books inner join Authors where Books.AuthorID = Authors.ID;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bookID, title, authorID, firstName, lastName string
		err := rows.Scan(&bookID, &title, &authorID, &firstName, &lastName)
		if err != nil {
			return nil, err
		}

		book := &model.Book{
			ID:    bookID,
			Title: title,
			Author: &model.Author{
				ID:        authorID,
				FirstName: firstName,
				LastName:  lastName,
			},
		}
		books = append(books, book)
	}

	return books, nil
}
