package taskstore

import "sync"

type Task struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type TaskStore struct {
	sync.Mutex

	tasks  map[int]Task
	nextId int
}

func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}

func (ts *TaskStore) GetAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	allTasks := make([]Task, 0, len(ts.tasks))
	for _, task := range ts.tasks {
		allTasks = append(allTasks, task)
	}
	return allTasks
}

func (ts *TaskStore) CreateTask(text string) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		Id:   ts.nextId,
		Text: text,
	}

	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task.Id
}
