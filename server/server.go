package server

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	storage "github.com/s-vvardenfell/BooksStorage/books_storage"
	"github.com/s-vvardenfell/BooksStorage/mysql_adapter"
	"github.com/sirupsen/logrus"
)

const (
	booksQuery = `SELECT title FROM books WHERE author_id = 
	(SELECT author_id FROM authors WHERE name = ?)`
	authorsQuery = `SELECT name FROM authors WHERE author_id IN 
	(SELECT author_id FROM books WHERE title = ?)`
)

// Сервер приложения и grpc-сервер
type Server struct {
	storage.UnimplementedBooksStorageServer
	ad *mysql_adapter.MySqlAdapter
}

// Создает и инизиализирует новый экземпляр Server
func New(driver, dsn string) *Server {
	ad, err := mysql_adapter.NewMySqlAdapter(driver, dsn)
	if err != nil {
		logrus.Error("cannot create server instance, %v", err)
	}
	return &Server{
		ad: ad,
	}
}

// Обрабатывает запрос получения списка книг по имени автора
func (s *Server) GetBooksByAuthor(
	ctx context.Context, in *storage.Author) (*storage.Books, error) {

	var err error
	res := &storage.Books{}

	// делаем запрос в бд
	if res.BookNames, err = s.ad.MakeQuery(booksQuery, in.AuthorName); err != nil {
		logrus.Errorf("error during request processing, %v", err)
		return nil, err
	}
	return res, nil
}

// Обрабатывает запрос получения списка авторов по названию книги
func (s *Server) GetAuthorsByBook(
	ctx context.Context, in *storage.Book) (*storage.Authors, error) {

	var err error
	res := &storage.Authors{}

	// делаем запрос в бд
	if res.AuthorNames, err = s.ad.MakeQuery(authorsQuery, in.BookName); err != nil {
		logrus.Errorf("error during request processing, %v", err)
		return nil, err
	}
	return res, nil
}

// Закрывает соединение с бд
func (s *Server) Stop() {
	_ = s.ad.Close()
}
