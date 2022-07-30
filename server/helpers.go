package server

// _ "github.com/go-sql-driver/mysql"

// // создает таблицу, если она не создана ранее
// func (s *Server) CreateTables(driver, dsn string) {

// 	stmtCreateAuthors, err := s.cnxn.Query(
// 		"CREATE TABLE IF NOT EXISTS authors (author_id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100))")
// 	if err != nil {
// 		logrus.Errorf("failed create table 'authors', %v", err)
// 	}
// 	defer stmtCreateAuthors.Close()

// 	stmtCreateBooks, err := s.cnxn.Query(
// 		`CREATE TABLE IF NOT EXISTS books (
// 			book_id INT PRIMARY KEY AUTO_INCREMENT,
// 			title VARCHAR(100),
// 			author_id INT,
// 			FOREIGN KEY (author_id)  REFERENCES authors (author_id)
// 			ON DELETE CASCADE)`)
// 	if err != nil {
// 		logrus.Errorf("failed create table 'books', %v", err)
// 	}
// 	defer stmtCreateBooks.Close()
// }

// // заполняет таблицу тестовыми данными
// func (s *Server) PopulateTables(driver, dsn string) {

// 	stmtInsAuthors, err := s.cnxn.Query(
// 		`INSERT INTO authors (name) VALUES
// 		('Пушкин А.С.'),
// 		('Достаевский Ф.М.'),
// 		('Чехов А.П.'),
// 		('Куприн А.И.'),
// 		('Стругацкий А.Н.'),
// 		('Стругацкий Б.Н.');`)
// 	if err != nil {
// 		logrus.Errorf("failed to populate table 'authors', %v", err)
// 	}
// 	defer stmtInsAuthors.Close()

// 	stmtInsBooks, err := s.cnxn.Query(
// 		`INSERT INTO books (title, author_id) VALUES
// 		('Капитанская дочка', (SELECT author_id FROM authors WHERE name = 'Пушкин А.С.')),
// 		('Идиот', (SELECT author_id FROM authors WHERE name = 'Достаевский Ф.М.')),
// 		('Вишневый сад', (SELECT author_id FROM authors WHERE name = 'Чехов А.П.')),
// 		('Гранатовый браслет', (SELECT author_id FROM authors WHERE name = 'Куприн А.И.')),
// 		('Понедельник начинается в субботу', (SELECT author_id FROM authors WHERE name = 'Стругацкий А.Н.')),
// 		('Понедельник начинается в субботу', (SELECT author_id FROM authors WHERE name = 'Стругацкий Б.Н.'));`)
// 	if err != nil {
// 		logrus.Errorf("failed to populate table 'books', %v", err)
// 	}
// 	defer stmtInsBooks.Close()
// }
