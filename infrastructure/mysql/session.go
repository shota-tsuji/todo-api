package mysql

import (
	"database/sql"
	"example.com/go-gin-todolist/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func NewMysqlSession(config config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@/%s", config.Mysql.User, config.Mysql.Password, config.Mysql.Database)
	db, err := sql.Open("mysql", dsn)
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
