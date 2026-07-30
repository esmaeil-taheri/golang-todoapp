package task

import (
	"errors"
	"fmt"
	"strings"

	"todoapp/entity"
	"todoapp/param"
)

type Repository interface {
	Create(t entity.Task) (entity.Task, error)
	GetByUserID(userID uint) ([]entity.Task, error)
}

// CategoryValidator lets the task service confirm a category belongs to the
// user without depending on the category service directly — it depends only on
// this narrow port. In practice the memory category store satisfies it.
type CategoryValidator interface {
	IsOwnedByUser(categoryID, userID uint) (bool, error)
}

type Service struct {
	repo Repository
	categories CategoryValidator
}

func New(repo Repository, categories CategoryValidator) Service {
	return Service{repo: repo, categories: categories}
}

func (s Service) Create(req param.CreateTaskRequest) (param.CreateTaskResponse, error) {
	if strings.TrimSpace(req.Title) == "" {
		return param.CreateTaskResponse{}, errors.New("task title can't be empty")
	}

	owned, err := s.categories.IsOwnedByUser(req.CategoryID, req.UserID)
	if err != nil {
		return param.CreateTaskResponse{}, fmt.Errorf("can't validate category: %w", err)
	}
	if !owned {
		return param.CreateTaskResponse{}, fmt.Errorf("category %d not found", req.CategoryID)
	}

	created, err := s.repo.Create(entity.Task{
		Title: req.Title,
		DueDate: req.DueDate,
		CategoryID: req.CategoryID,
		IsDone: false,
		UserID: req.UserID,
	})
	if err != nil {
		return param.CreateTaskResponse{}, fmt.Errorf("can't create task: %w", err)
	}

	return param.CreateTaskResponse{Task: created}, nil
}

func (s Service) List(req param.ListTaskRequest) (param.ListTaskResponse, error) {
	tasks, err := s.repo.GetByUserID(req.UserID)
	if err != nil {
		return param.ListTaskResponse{}, fmt.Errorf("can't list tasks: %w", err)
	}

	return param.ListTaskResponse{Tasks: tasks}, nil
}
