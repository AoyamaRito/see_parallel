package queue

import (
	"sync"
)

type Task struct {
	Question string
	Files    []string
	Deep     bool
}

type Queue struct {
	tasks []Task
	mu    sync.Mutex
}

var instance *Queue
var once sync.Once

func GetInstance() *Queue {
	once.Do(func() {
		instance = &Queue{
			tasks: make([]Task, 0),
		}
	})
	return instance
}

func (q *Queue) Add(task Task) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.tasks = append(q.tasks, task)
}

func (q *Queue) GetAll() []Task {
	q.mu.Lock()
	defer q.mu.Unlock()
	tasksCopy := make([]Task, len(q.tasks))
	copy(tasksCopy, q.tasks)
	return tasksCopy
}

func (q *Queue) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.tasks = make([]Task, 0)
}

func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.tasks)
}