package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := sql.DB{}
	tx, _ := db.Begin()
	tx.Commit()
	defer tx.Rollback()

}
