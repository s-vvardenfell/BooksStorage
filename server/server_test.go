package server

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	storage "github.com/s-vvardenfell/BooksStorage/books_storage"
)

func RunServer(driver, dsn, host, port string) {
	grpcServ := grpc.NewServer()
	serv := New(driver, dsn)
	storage.RegisterBooksStorageServer(grpcServ, serv)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	logrus.Info("Starting gRPC listener on port " + port)
	if err := grpcServ.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func CreateTables(driver, dsn string) {
	cnxn, err := sql.Open(driver, dsn)
	if err != nil {
		logrus.Fatalf("failed to open connection, %v", err)
	}
	defer cnxn.Close()

	stmt, err := cnxn.Query(
		"CREATE TABLE IF NOT EXISTS authors (author_id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100))")
	if err != nil {
		logrus.Fatalf("failed to prepare statement, %v", err)
	}
	defer stmt.Close()

	stmt, err = cnxn.Query(
		`CREATE TABLE IF NOT EXISTS books (
			book_id INT PRIMARY KEY AUTO_INCREMENT, 
			title VARCHAR(100), 
			author_id INT, 
			FOREIGN KEY (author_id)  REFERENCES authors (author_id)
			ON DELETE CASCADE)`)
	if err != nil {
		logrus.Fatalf("failed to prepare statement, %v", err)
	}
	defer stmt.Close()
}

func PopulateTables(driver, dsn string) {
	cnxn, err := sql.Open(driver, dsn)
	if err != nil {
		logrus.Fatalf("failed to open connection, %v", err)
	}
	defer cnxn.Close()

	stmt, err := cnxn.Query(
		`INSERT INTO authors (name) VALUES
		('Пушкин А.С.'), 
		('Достаевский Ф.М.'), 
		('Чехов А.П.'),
		('Куприн А.И.'), 
		('Стругацкий А.Н.'), 
		('Стругацкий Б.Н.');`)
	if err != nil {
		logrus.Fatalf("failed to populate authors table, %v", err)
	}
	defer stmt.Close()

	stmt, err = cnxn.Query(
		`INSERT INTO books (title, author_id) VALUES
		('Капитанская дочка', (SELECT author_id FROM authors WHERE name = 'Пушкин А.С.')), 
		('Идиот', (SELECT author_id FROM authors WHERE name = 'Достаевский Ф.М.')), 
		('Вишневый сад', (SELECT author_id FROM authors WHERE name = 'Чехов А.П.')), 
		('Гранатовый браслет', (SELECT author_id FROM authors WHERE name = 'Куприн А.И.')), 
		('Понедельник начинается в субботу', (SELECT author_id FROM authors WHERE name = 'Стругацкий А.Н.')), 
		('Понедельник начинается в субботу', (SELECT author_id FROM authors WHERE name = 'Стругацкий Б.Н.'));`)
	if err != nil {
		logrus.Fatalf("failed to populate books table, %v", err)
	}
	defer stmt.Close()
}

func Test_pkg(t *testing.T) {
	// CreateTables("mysql", "")
	// PopulateTables("mysql", "")

	go RunServer("mysql", "s.vvardenfell:Zxasqw12@/kvadoru", "localhost", "50052")
	time.Sleep(time.Second)

	client := NewClientStub("localhost", "50052")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Log("\tGetting books by author name")
	{
		author := "Достаевский Ф.М."
		res, err := client.GetBooksByAuthor(ctx, &storage.Author{AuthorName: author})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, res.BookNames[0], "Идиот")
	}

	t.Log("\tGetting author by book title")
	{
		book := "Идиот"
		res, err := client.GetAuthorsByBook(ctx, &storage.Book{BookName: book})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, res.AuthorNames[0], "Достаевский Ф.М.")
	}

	t.Log("\tGetting book authors (more that 1 author)")
	{
		book := "Понедельник начинается в субботу"
		res, err := client.GetAuthorsByBook(ctx, &storage.Book{BookName: book})
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, len(res.AuthorNames), 2)
	}

	t.Log("\tGetting books by wrong author name")
	{
		author := "Толкин Д.Р.Р."
		res, err := client.GetBooksByAuthor(ctx, &storage.Author{AuthorName: author})
		require.NoError(t, err)
		require.Empty(t, res.BookNames)
	}

	t.Log("\tGetting author by wrong book title")
	{
		book := "Властелин колец"
		res, err := client.GetAuthorsByBook(ctx, &storage.Book{BookName: book})
		require.NoError(t, err)
		require.Empty(t, res.AuthorNames)
	}
}
