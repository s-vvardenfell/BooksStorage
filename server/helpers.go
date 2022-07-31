package server

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// _ "github.com/go-sql-driver/mysql"

// создает таблицу, если она не создана ранее
func CreateTables(cnxn *sql.DB) error {

	stmtCreateAuthors, err := cnxn.Query(
		"CREATE TABLE IF NOT EXISTS authors (author_id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100))")
	if err != nil {
		logrus.Errorf("failed create table 'authors', %v", err)
		return err
	}
	defer stmtCreateAuthors.Close()

	stmtCreateBooks, err := cnxn.Query(
		`CREATE TABLE IF NOT EXISTS books (
			book_id INT PRIMARY KEY AUTO_INCREMENT,
			title VARCHAR(100),
			author_id INT,
			FOREIGN KEY (author_id)  REFERENCES authors (author_id)
			ON DELETE CASCADE)`)
	if err != nil {
		logrus.Errorf("failed create table 'books', %v", err)
		return err
	}
	defer stmtCreateBooks.Close()
	return nil
}

// заполняет таблицу тестовыми данными
func PopulateTables(cnxn *sql.DB) error {

	stmtInsAuthors, err := cnxn.Query(
		`INSERT INTO authors (name) VALUES
		('Пушкин А.С.'),
		('Достаевский Ф.М.'),
		('Чехов А.П.'),
		('Куприн А.И.'),
		('Стругацкий А.Н.'),
		('Стругацкий Б.Н.');`)
	if err != nil {
		logrus.Errorf("failed to populate table 'authors', %v", err)
		return err
	}
	defer stmtInsAuthors.Close()

	stmtInsBooks, err := cnxn.Query(
		`INSERT INTO books (title, author_id) VALUES
		('Капитанская дочка', (SELECT author_id FROM authors WHERE name = 'Пушкин А.С.')),
		('Идиот', (SELECT author_id FROM authors WHERE name = 'Достаевский Ф.М.')),
		('Преступление и наказание', (SELECT author_id FROM authors WHERE name = 'Достаевский Ф.М.')),
		('Вишневый сад', (SELECT author_id FROM authors WHERE name = 'Чехов А.П.')),
		('Гранатовый браслет', (SELECT author_id FROM authors WHERE name = 'Куприн А.И.')),
		('Понедельник начинается в субботу', (SELECT author_id FROM authors WHERE name = 'Стругацкий А.Н.')),
		('Понедельник начинается в субботу', (SELECT author_id FROM authors WHERE name = 'Стругацкий Б.Н.'));`)
	if err != nil {
		logrus.Errorf("failed to populate table 'books', %v", err)
		return err
	}
	defer stmtInsBooks.Close()
	return nil
}

func CleanUp(cnxn *sql.DB) error {

	dropBooks, err := cnxn.Query(
		`DROP TABLE books`)
	if err != nil {
		logrus.Errorf("failed to drop table 'books', %v", err)
		return err
	}
	defer dropBooks.Close()

	dropAuthors, err := cnxn.Query(
		`DROP TABLE authors`)
	if err != nil {
		logrus.Errorf("failed to drop table 'authors', %v", err)
		return err
	}
	defer dropAuthors.Close()

	return nil
}
