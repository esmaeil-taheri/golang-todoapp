package filestore

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"todoapp/entity"
)


type UserStore struct {
	mu sync.Mutex
	filePath string
	users []entity.User
	lastID uint
}


func NewUserStore(filePath string) (*UserStore, error) {
	s := &UserStore{filePath: filePath, users: make([]entity.User, 0)}
	if err := s.load(); err != nil {
		return nil, err
	}

	return s, nil
}


func (s *UserStore) load() error {
	file, err := os.Open(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// No file yet: start empty, it's created on the first save.
			return nil
		}

		return fmt.Errorf("can't open user file: %w", err)
	}
	file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		var u entity.User
		if err := json.Unmarshal(line, &u); err != nil {
			return fmt.Errorf("can't decode user record: %w", err)
		}

		s.users = append(s.users, u)
		if u.ID > s.lastID {
			s.lastID = u.ID
		}
	}

	return scanner.Err()
}

func (s *UserStore) Create(u entity.User) (entity.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++
	u.ID = s.lastID

	if err := s.appnedToFile(u); err != nil {
		s.lastID--
		return entity.User{}, err
	}
	s.users = append(s.users, u)

	return u, nil
}


func (s *UserStore) GetByEmail(email string) (entity.User, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, u := range s.users {
		if u.Email == email{
			return u, true, nil
		}
	}

	return entity.User{}, false, nil
}


func (s *UserStore) appnedToFile(u entity.User) error {
	file, err := os.OpenFile(s.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return fmt.Errorf("can't open user file for writing: %w", err)
	}
	defer file.Close()

	data, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("can't encode user: %w", err)
	}

	data = append(data, '\n')

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("can't write user: %w", err)
	}

	return nil

}