package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	v1 "github.com/sazid/learngrpc/api/v1"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save(*v1.Laptop) error
}

type InMemoryLaptopStore struct {
	sync.RWMutex
	data map[string]*v1.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*v1.Laptop),
	}
}

func (s *InMemoryLaptopStore) Save(laptop *v1.Laptop) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.data[laptop.Id]; !ok {
		return ErrAlreadyExists
	}

	other := &v1.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	s.data[other.Id] = other

	return nil
}
