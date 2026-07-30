package category

import (
	"errors"
	"strings"

	"todoapp/entity"
	"todoapp/param"
)

type Repository interface {
	Create(c entity.Category) (entity.Category, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Create(req param.CreateCategoryRequest) (param.CreateCategoryResponse, error) {
	if strings.TrimSpace(req.Title) == "" {
		return param.CreateCategoryResponse{}, errors.New("category title can't be empty")
	}

	created, err := s.repo.Create(entity.Category{
		Title: req.Title,
		Color: req.Color,
		UserID: req.UserID,
	})
	if err != nil {
		return param.CreateCategoryResponse{}, err
	}

	return param.CreateCategoryResponse{Category: created}, nil
}
