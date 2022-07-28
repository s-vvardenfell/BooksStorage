package server

import (
	"context"
	"database/sql"

	storage "github.com/s-vvardenfell/BooksStorage/books_storage"
	"github.com/sirupsen/logrus"
)

type Server struct {
	storage.UnimplementedBooksStorageServer
	db *sql.DB
}

func New(driver, dsn string) *Server {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		logrus.Fatalf("cannot open %s connection, %v", driver, err)
	}

	return &Server{
		db: db,
	}
}

func (s *Server) GetBooksByAuthors(
	ctx context.Context, in *storage.Authors) (*storage.Books, error) {

	err := s.db.Ping()
	if err != nil {
		//TODO RE-OPEN DB (если не откроется снова, слать ошибку)
		logrus.Errorf("cannot connect with DataBase, %v", err)
		return nil, err //мб какое-то описание сделать
	}

	stmt, err := s.db.Prepare(
		`SELECT title FROM books WHERE author_id = 
		(SELECT author_id FROM authors WHERE name = ?)`)
	if err != nil {
		logrus.Errorf("cannot prepare statement, %v", err)
		return nil, err //мб какое-то описание сделать
	}
	defer stmt.Close() //TODO лямбда, обработать возврат

	res := &storage.Books{}
	res.BookName = make([]string, 0)

	//TODO здесь будет 1 арг
	for _, name := range in.AuthorName {
		row, err := stmt.Query(name)
		if err != nil {
			logrus.Errorf("cannot make query, %v", err)
			return nil, err //мб какое-то описание сделать
		}

		for row.Next() {
			var temp string
			row.Scan(&temp)
			res.BookName = append(res.BookName, temp)
		}
	}

	return res, nil
}

func (s *Server) GetAuthorsByBooks(
	ctx context.Context, in *storage.Books) (*storage.Authors, error) {

	err := s.db.Ping()
	if err != nil {
		//TODO RE-OPEN DB (если не откроется снова, слать ошибку)
		return nil, err //мб какое-то описание сделать
	}

	stmt, err := s.db.Prepare(
		`SELECT name FROM authors INNER JOIN books 
		ON authors.author_id = books.author_id 
		WHERE books.author_id IN 
		(SELECT author_id FROM books WHERE title = ?)`)
	if err != nil {
		logrus.Errorf("cannot prepare statement, %v", err)
		return nil, err //мб какое-то описание сделать
	}
	defer stmt.Close() //TODO лямбда, обработать возврат

	res := &storage.Authors{}
	res.AuthorName = make([]string, 0)

	//TODO здесь будет 1 арг
	//TODO получить row и потом посчитать кол-во элементов, выделить память
	for _, name := range in.BookName {
		row, err := stmt.Query(name)
		if err != nil {
			logrus.Errorf("cannot make query, %v", err)
			return nil, err //мб какое-то описание сделать
		}

		for row.Next() {
			var temp string
			row.Scan(&temp)
			res.AuthorName = append(res.AuthorName, temp)
		}
	}

	return res, nil
}
