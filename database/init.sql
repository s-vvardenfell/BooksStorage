CREATE DATABASE IF NOT EXISTS kvadoru;

USE kvadoru;

CREATE TABLE IF NOT EXISTS authors (author_id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100));

INSERT INTO authors (name) VALUES
		('Пушкин А.С.'), 
		('Достаевский Ф.М.'), 
		('Чехов А.П.'),
		('Куприн А.И.'), 
		('Стругацкий А.Н.'), 
		('Стругацкий Б.Н.');

CREATE TABLE IF NOT EXISTS books (
      book_id INT PRIMARY KEY AUTO_INCREMENT, 
      title VARCHAR(100), 
      author_id INT, 
      FOREIGN KEY (author_id)  REFERENCES authors (author_id)
      ON DELETE CASCADE 
);
		
INSERT INTO books (title, author_id) VALUES
('Капитанская дочка', '1'), 
('Идиот', '2'), 
('Преступление и наказаение', '2'),
('Вишневый сад', '3'), 
('Гранатовый браслет', '4'), 
('Понедельник начинается в субботу', '5'), 
('Понедельник начинается в субботу', '6');
