package instance

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Pq() {
	db, err := sql.Open("postgres", "user=mdvis dbname=mydb sslmode=disable")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("laotie2", "dev", "2000-1-2")
	checkErr(err)

	fmt.Println(res)
}
