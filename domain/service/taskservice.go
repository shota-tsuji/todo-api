package service

import (
	"database/sql"
	"example.com/go-gin-todolist/domain/entity"
	"sync"
)

type TaskService struct {
	sync.Mutex

	mysqlSession *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{mysqlSession: db}
}

func (ts *TaskService) GetAllTasks() []entity.Task {
	ts.Lock()
	defer ts.Unlock()

	rows, _ := ts.mysqlSession.Query("select * from task")

	var allTasks []entity.Task
	for rows.Next() {
		var id int
		var title string
		rows.Scan(&id, &title)
		allTasks = append(allTasks, entity.Task{Id: id, Text: title})
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
