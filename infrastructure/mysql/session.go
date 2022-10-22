package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func NewMysqlSession() *sql.DB {
	db, err := sql.Open("mysql", "root:password@/todo")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		log.Fatal("fatal")
	} else {
		log.Println("success")
	}

	return db
}
