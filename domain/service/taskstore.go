package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"time"
)

type Task struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type TaskService struct {
	sync.Mutex

	mysqlSession *sql.DB
}

func New() *TaskService {
	ts := &TaskService{}

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

	ts.mysqlSession = db
	return ts
}

func (ts *TaskService) GetAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	rows, _ := ts.mysqlSession.Query("select * from task")

	var allTasks []Task
	for rows.Next() {
		var id int
		var title string
		rows.Scan(&id, &title)
		allTasks = append(allTasks, Task{Id: id, Text: title})
	}

	return allTasks
}

func (ts *TaskService) CreateTask(title string) int {
	ts.Lock()
	defer ts.Unlock()

	row := ts.mysqlSession.QueryRow("select count(*) from task")
	var count int
	row.Scan(&count)
	id := count + 1

	result, err := ts.mysqlSession.Exec("insert into task(id, title) values (?, ?)",
		id, title)
	_, err = result.LastInsertId()

	if err != nil {
		return -1
	}

	return id
}
