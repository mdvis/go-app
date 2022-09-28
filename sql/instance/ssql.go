package instance

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Sq() {
	db, err := sql.Open("sqlite3", "/Users/mdvis/foo.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname,created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("laotie", "dev", "2000-1-2")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}
