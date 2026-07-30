package user

import (
	"errors"
	"fmt"
	"strings"

	"todoapp/entity"
	"todoapp/param"
	"todoapp/pkg/password"
)

// Repository is the port this service needs for user persistence. Any adapter
// (file, memory, postgres...) that implements it can be plugged in without the
// service changing.
type Repository interface {
	Create(u entity.User) (entity.User, error)
	GetByEmail(email string) (entity.User, bool, error)
}

type Service struct {
	repo Repository
	bcryptCost int
}

func New(repo Repository, bcryptCost int) Service {
	return  Service{repo: repo, bcryptCost: bcryptCost}
}

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {
	if err := validateRegister(req); err != nil {
		return param.RegisterResponse{}, err
	}

	_, exists, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if exists {
		return param.RegisterResponse{}, errors.New("this email is already registered")
	}

	hashed, err := password.Hash(req.Password, s.bcryptCost)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("can't hash password: %w", err)
	}

	created, err := s.repo.Create(entity.User{
		Name: req.Name,
		Email: req.Email,
		Password: hashed,
	})
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("can't create user: %w", err)
	}

	return param.RegisterResponse{User: toUserInfo(created)}, nil
}

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	user, found, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	// Same generic message whether the email is unknown or the password is
	// wrong, so we don't leak which emails exist.
	if !found || !password.Compare(user.Password, req.Password) {
		return param.LoginResponse{}, errors.New("email or password is incorrect")
	}

	return param.LoginResponse{User: toUserInfo(user)}, nil
}

func validateRegister(req param.RegisterRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name can't be empty")
	}
	if !strings.Contains(req.Email, "@") {
		return errors.New("email is not valid")
	}
	if len(req.Password) < 4 {
		return errors.New("password must be at least 4 charecters")
	}

	return nil
}

func toUserInfo(u entity.User) param.UserInfo {
	return param.UserInfo{ID: u.ID, Name: u.Name, Email: u.Email}
}
