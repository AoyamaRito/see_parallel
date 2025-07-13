package queue

import (
	"encoding/json"
	"os"
	"path/filepath"
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
		instance.loadFromFile()
	})
	return instance
}

func (q *Queue) Add(task Task) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.tasks = append(q.tasks, task)
	q.saveToFile()
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
	q.saveToFile()
}

func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.tasks)
}

func (q *Queue) getQueueFile() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(wd, ".see_parallel", "queue.json")
}

func (q *Queue) saveToFile() {
	queueFile := q.getQueueFile()
	if queueFile == "" {
		return
	}

	dir := filepath.Dir(queueFile)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return
	}

	data, err := json.Marshal(q.tasks)
	if err != nil {
		return
	}

	os.WriteFile(queueFile, data, 0600)
}

func (q *Queue) loadFromFile() {
	queueFile := q.getQueueFile()
	if queueFile == "" {
		return
	}

	data, err := os.ReadFile(queueFile)
	if err != nil {
		return
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return
	}

	q.tasks = tasks
}