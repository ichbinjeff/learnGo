package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=shli15 dbname=sample sslmode=disable")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo (username, departname, created) VALUES ($1, $2, $3) RETURNING uid;")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "R&D", "2017-02-11")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username, departname, created) VALUES ($1, $2, $3) RETURNING UID;", "astaxie2", "R&D", "2017-02-13").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("最后插入的id＝", lastInsertId)

	// Update data
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", "1")
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// retrieve
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(1)
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
