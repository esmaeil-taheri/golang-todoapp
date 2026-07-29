// Package memory holds in-RAM repository adapters. Each store guards its state
// with a mutex, so it stays correct once the HTTP server serves every request
// in its own goroutine.
package memory

import (
	"sync"

	"todoapp/entity"
)

type TaskStore struct {
	mu sync.Mutex
	tasks []entity.Task
	lastID uint
}

func NewTaskStore() *TaskStore {
	return &TaskStore{tasks: make([]entity.Task, 0)}
}

func (s *TaskStore) Create(t entity.Task) (entity.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++
	t.ID = s.lastID
	s.tasks = append(s.tasks, t)

	return t, nil
}

func (s * TaskStore) GetByUserID(userID uint) ([]entity.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]entity.Task, 0)
	for _, t := range s.tasks {
		if t.UserID == userID {
			result = append(result, t)
		}
	}

	return result, nil
}