package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db_connect_timeout = 200
	db_read_timeout    = 1000
	db_write_timeout   = 1000
	db_username        = "xudahai"
	db_password        = "test123"
	db_dbname          = "notify"
	db_charset         = "utf8"
	db_hosts           = "192.168.240.84:3306"
)

func main() {
	db, err := sql.Open("mysql", "nichao:nichao@tcp(127.0.0.1:3306)/notify?charset=utf8")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		insertSql := "insert into  tblOnlineStat0 (uid, product, extra, cmdno, stime, utime, dt ) values (?, ?, ?, ?, ? ,?, ?) "
		if smt, err := db.Prepare(insertSql); err == nil {
			if _, err := smt.Exec(1, "fudao", "uid=1", 100011, 1501234567, 1501234567, 20170818); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("success")
			}
		} else {
			fmt.Println(err)
		}

	}
}
