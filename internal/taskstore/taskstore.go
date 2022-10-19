package taskstore

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

type TaskStore struct {
	sync.Mutex

	db *sql.DB
}

func New() *TaskStore {
	ts := &TaskStore{}

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

	ts.db = db
	return ts
}

func (ts *TaskStore) GetAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	rows, _ := ts.db.Query("select * from task")

	var allTasks []Task
	for rows.Next() {
		var id int
		var title string
		rows.Scan(&id, &title)
		allTasks = append(allTasks, Task{Id: id, Text: title})
	}

	return allTasks
}

func (ts *TaskStore) CreateTask(title string) int {
	ts.Lock()
	defer ts.Unlock()

	row := ts.db.QueryRow("select count(*) from task")
	var count int
	row.Scan(&count)
	id := count + 1

	result, err := ts.db.Exec("insert into task(id, title) values (?, ?)",
		id, title)
	_, err = result.LastInsertId()

	if err != nil {
		return -1
	}

	return id
}
