package mysql

import (
	"database/sql"
	"example.com/go-gin-todolist/domain/entity"
	"sync"
)

type Repository struct {
	sync.Mutex

	mysqlSession *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{mysqlSession: db}
}

func (r *Repository) InsertTask(title string) int {
	r.Lock()
	defer r.Unlock()

	row := r.mysqlSession.QueryRow("select count(*) from task")
	var count int
	row.Scan(&count)
	id := count + 1

	result, err := r.mysqlSession.Exec("insert into task(id, title) values (?, ?)",
		id, title)
	_, err = result.LastInsertId()

	if err != nil {
		return -1
	}

	return id
}

func (r *Repository) FindAllTasks() []entity.Task {
	r.Lock()
	defer r.Unlock()

	rows, _ := r.mysqlSession.Query("select * from task")

	var allTasks []entity.Task
	for rows.Next() {
		var id int
		var title string
		rows.Scan(&id, &title)
		allTasks = append(allTasks, entity.Task{Id: id, Text: title})
	}

	return allTasks
}
