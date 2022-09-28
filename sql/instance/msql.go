package instance

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Mq() {
	db, err := sql.Open("mysql", "root:123456@/mydb?charset=utf8")
	checkErr(err)

	// Insert
	stmt, err := db.Prepare("INSERT userinfo SET USERNAME=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("root", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// Update
	stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec("rootupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// Search
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var (
			uid        int
			username   string
			department string
			created    string
		)
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println(uid, username, department, created)
	}

	// Delete
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
