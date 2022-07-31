package mysql_adapter

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func Test_MakeQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	t.Log("Getting books sub-query test")
	{
		columns := []string{"2"}
		mock.ExpectPrepare("SELECT author_id FROM authors WHERE name = (.+)").
			ExpectQuery().
			WithArgs("Достаевский Ф.М.").
			WillReturnRows(sqlmock.NewRows(columns))

		var ad MySqlAdapter
		ad.cnxn = db

		_, err = ad.MakeQuery("SELECT author_id FROM authors WHERE name = ?",
			"Достаевский Ф.М.")
		require.NoError(t, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err)
	}

	t.Log("Getting books query test")
	{
		columns := []string{"Идиот", "Преступление и наказание"}
		mock.ExpectPrepare("SELECT title FROM books WHERE author_id = (.+)").
			ExpectQuery().
			WithArgs("2").
			WillReturnRows(sqlmock.NewRows(columns))

		var ad MySqlAdapter
		ad.cnxn = db

		_, err = ad.MakeQuery("SELECT title FROM books WHERE author_id = ?", "2")
		require.NoError(t, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err)
	}

	t.Log("Getting author sub-query test")
	{
		columns := []string{"5"}
		mock.ExpectPrepare("SELECT author_id FROM books WHERE title = (.+)").
			ExpectQuery().
			WithArgs("Понедельник начинается в субботу").
			WillReturnRows(sqlmock.NewRows(columns))

		var ad MySqlAdapter
		ad.cnxn = db

		_, err = ad.MakeQuery("SELECT author_id FROM books WHERE title = ?",
			"Понедельник начинается в субботу")
		require.NoError(t, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err)
	}

	t.Log("Getting author query test")
	{
		columns := []string{"Стругацкий Б.Н.", "Стругацкий А.Н."}
		mock.ExpectPrepare("SELECT name FROM authors WHERE author_id IN (.+)").
			ExpectQuery().
			WithArgs("5").
			WillReturnRows(sqlmock.NewRows(columns))

		var ad MySqlAdapter
		ad.cnxn = db

		_, err = ad.MakeQuery("SELECT name FROM authors WHERE author_id IN ?", "5")
		require.NoError(t, err)

		err = mock.ExpectationsWereMet()
		require.NoError(t, err)
	}
}
