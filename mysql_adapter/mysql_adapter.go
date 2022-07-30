package mysql_adapter

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

const (
	connMaxLiveTime = 3
	maxOpenCnxns    = 10
	maxIdleCnxns    = 10
	maxAttempts     = 3
)

// Cодержит соединение с бд и выполняет операции обращения к бд
type MySqlAdapter struct {
	cnxn *sql.DB
}

func NewMySqlAdapter(driver, dsn string) (*MySqlAdapter, error) {
	cnxn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	cnxn.SetConnMaxLifetime(time.Minute * connMaxLiveTime)
	cnxn.SetMaxOpenConns(maxOpenCnxns)
	cnxn.SetMaxIdleConns(maxIdleCnxns)

	// проверяем соединение с бд
	attempts := 0
	for {
		logrus.Infof("Trying to connect to database, attempt #%d", attempts)

		if err := cnxn.Ping(); err != nil {
			logrus.Errorf("Failed to connect to database")
			attempts++
			if attempts == maxAttempts {
				return nil, err
			}
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	return &MySqlAdapter{cnxn: cnxn}, nil
}

func (s *MySqlAdapter) Close() error {
	return s.cnxn.Close()
}

// MakeQuery делает запрос к базе данных, на вход принимает
// sql-запрос и значение для подстановки,
// возвращает слайс строк или ошибку в случае неудачи
func (s *MySqlAdapter) MakeQuery(query, val string) ([]string, error) {
	// готовим запрос
	stmt, err := s.cnxn.Prepare(query)
	if err != nil {
		logrus.Errorf("cannot prepare statement, %v", err)
		return nil, err
	}
	defer stmt.Close()

	row, err := stmt.Query(val)
	if err != nil {
		logrus.Errorf("cannot make query, %v", err)
		return nil, err
	}

	res := make([]string, 0)

	for row.Next() {
		var temp string
		if err := row.Scan(&temp); err != nil {
			logrus.Errorf("cannot scan value, %v", err)
			return nil, err
		}
		res = append(res, temp)
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
