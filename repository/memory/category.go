package memory

import (
	"sync"

	"todoapp/entity"
)

type CategoryStore struct {
	mu sync.Mutex
	categories []entity.Category
	lastID uint
}

func NewCategoryStore() *CategoryStore {
	return &CategoryStore{categories: make([]entity.Category, 0)}
}

func (s *CategoryStore) Create(c entity.Category) (entity.Category, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID ++
	c.ID = s.lastID
	s.categories = append(s.categories, c)

	return c, nil
}

// IsOwnedByUser satisfies task.CategoryValidator: one adapter can implement
// several ports, exactly like a driven adapter in hexagonal architecture.
func (s *CategoryStore) IsOwnedByUser(categoryID, userID uint) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, c := range s.categories {
		if c.ID == categoryID && c.UserID == userID {
			return true, nil
		}
	}

	return false, nil
}