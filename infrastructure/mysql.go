package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/ntrianta/cerndemo/interfaces"
)

type MySQL struct {
	Db *sql.DB
}

type Rows struct {
	Rows *sql.Rows
}

//MySQL

func (mysql MySQL) Store(query string, writable ...interface{}) error {

	fmt.Println("Infrastructure layer. MySQL Handler. Storage")

	prep, _ := mysql.Db.Prepare(query)
	_, _ = prep.Exec(writable...) //Fix error handling
	return nil
}

func (mysql MySQL) List(query string) interfaces.Rows {
	rows, _ := mysql.Db.Query(query)
	return rows
}

func (mysql MySQL) Read(query string, id int, readable *string) {

	fmt.Println("Infrastructure layer. MySQL Handler. Read")

	row := mysql.Db.QueryRow(query, id)
	_ = row.Scan(readable)
}

func (mysql MySQL) Delete(query string, id int) error {
	prep, _ := mysql.Db.Prepare(query)
	_, _ = prep.Exec(id)
	return nil
}

//Rows

func (rows Rows) Scan(readable ...interface{}) error {

	rows.Rows.Scan(readable)
	return nil
}

func (rows Rows) Next() bool {

	return rows.Rows.Next()
}

func (rows Rows) Close() error {

	rows.Rows.Close()
	return nil
}
