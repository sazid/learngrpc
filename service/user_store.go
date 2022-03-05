package service

import (
	"fmt"
	"sync"
)

// UserStore is an interface to store and interact with users
type UserStore interface {
	// Save saves a new user to the store
	Save(user *User) error
	// Find finds a user by username
	Find(username string) (*User, error)
}

// InMemoryUserStore stores users in memory
type InMemoryUserStore struct {
	sync.RWMutex
	users map[string]*User
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*User),
	}
}

func (s *InMemoryUserStore) Save(user *User) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.users[user.Username]; ok {
		return ErrAlreadyExists
	}

	s.users[user.Username] = user.Clone()

	return nil
}

func (s *InMemoryUserStore) Find(username string) (*User, error) {
	s.RLock()
	defer s.RUnlock()

	if user, ok := s.users[username]; ok {
		return user.Clone(), nil
	}

	return nil, fmt.Errorf("cannot find user with the username: %s", username)
}
