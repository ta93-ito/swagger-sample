package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	const sourceName = "root:password@tcp(mysql:3306)/swagger_sample?charset=utf8&parseTime=True&Local"
	var err error
	DB, err = sql.Open("mysql", sourceName)
	if err != nil {
		panic(err.Error())
	}
}
