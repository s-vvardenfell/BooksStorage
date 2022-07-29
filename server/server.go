package server

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	storage "github.com/s-vvardenfell/BooksStorage/books_storage"
	"github.com/sirupsen/logrus"
)

const (
	countBooksQuery = `SELECT COUNT(title) FROM books WHERE author_id = 
	(SELECT author_id FROM authors WHERE name = ?)`
	booksQuery = `SELECT title FROM books WHERE author_id = 
	(SELECT author_id FROM authors WHERE name = ?)`
	countAuthorsQuery = `SELECT COUNT(name) FROM authors WHERE author_id IN 
	(SELECT author_id FROM books WHERE title = ?)`
	authorsQuery = `SELECT name FROM authors WHERE author_id IN 
	(SELECT author_id FROM books WHERE title = ?)`
)

type Server struct {
	storage.UnimplementedBooksStorageServer
	cnxn   *sql.DB
	driver string
	dsn    string
}

func New(driver, dsn string) *Server {
	serv := Server{
		driver: driver,
		dsn:    dsn,
	}

	if err := serv.openDbConnection(); err != nil {
		logrus.Fatalf("cannot open %s connection, %v", driver, err)
	}
	return &serv
}

func (s *Server) openDbConnection() (err error) {
	s.cnxn, err = sql.Open(s.driver, s.dsn)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) GetBooksByAuthor(
	ctx context.Context, in *storage.Author) (*storage.Books, error) {

	// проверяем соединение с бд
	if err := s.cnxn.Ping(); err != nil {
		// если ошибка соединения, пытаемся установить новое
		if err := s.openDbConnection(); err != nil {
			logrus.Errorf("cannot re-open %s connection, %v", s.driver, err)
			return nil, err
		}
	}

	// для того, чтобы выделить достаточно памяти
	// сразу посчитаем кол-во результатов
	resultNum, err := s.getResultsCount(countBooksQuery, in.AuthorName)
	if err != nil {
		logrus.Errorf("cannot count results number, %v", err)
		return nil, err
	}

	// выделяем память
	res := &storage.Books{}
	res.BookNames = make([]string, 0, resultNum)

	// делаем запрос в бд
	if err := s.makeQuery(&res.BookNames, booksQuery, in.AuthorName); err != nil {
		logrus.Errorf("error during request processing, %v", err)
		return nil, err
	}
	return res, nil
}

func (s *Server) GetAuthorsByBook(
	ctx context.Context, in *storage.Book) (*storage.Authors, error) {

	// проверяем соединение с бд
	if err := s.cnxn.Ping(); err != nil {
		// если ошибка соединения, пытаемся установить новое
		if err := s.openDbConnection(); err != nil {
			logrus.Errorf("cannot re-open %s connection, %v", s.driver, err)
			return nil, err
		}
	}

	// для того, чтобы выделить достаточно памяти
	// сразу посчитаем кол-во результатов
	resultNum, err := s.getResultsCount(countAuthorsQuery, in.BookName)
	if err != nil {
		logrus.Errorf("cannot count results number, %v", err)
		return nil, err
	}

	// выделяем память
	res := &storage.Authors{}
	res.AuthorNames = make([]string, 0, resultNum)

	// делаем запрос в бд
	if err := s.makeQuery(&res.AuthorNames, authorsQuery, in.BookName); err != nil {
		logrus.Errorf("error during request processing, %v", err)
		return nil, err
	}
	return res, nil
}

// считает количество результатов для выборки,
// возвращает число и/или ошибку в случае неудачи
func (s *Server) getResultsCount(countQuery, prepValue string) (int, error) {
	var resultNum int
	// готовим запрос
	stmt, err := s.cnxn.Prepare(countQuery)
	if err != nil {
		logrus.Errorf("cannot prepare statement, %v", err)
		return resultNum, err
	}
	defer stmt.Close()

	// получаем результаты
	err = stmt.QueryRow(prepValue).Scan(&resultNum)
	if err != nil {
		logrus.Errorf("cannot make query, %v", err)
		return resultNum, err
	}
	return resultNum, nil
}

// делает запрос к базе данных,
// на вход принимает слайс для хранения результатов,
// sql-запрос и аргумент для prepared statement,
// возвращает ошибку в случае неудачи
func (s *Server) makeQuery(res *[]string, query, prepValue string) error {
	// готовим запрос
	stmt, err := s.cnxn.Prepare(query)
	if err != nil {
		logrus.Errorf("cannot prepare statement, %v", err)
		return err
	}
	defer stmt.Close()

	// получаем результаты
	row, err := stmt.Query(prepValue)
	if err != nil {
		logrus.Errorf("cannot make query, %v", err)
		return err
	}

	// извлекаем
	for row.Next() {
		var temp string
		if err := row.Scan(&temp); err != nil {
			logrus.Errorf("cannot scan value, %v", err)
			return err
		}
		*res = append(*res, temp)
	}
	return nil
}
