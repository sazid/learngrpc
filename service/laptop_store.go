package service

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/copier"
	v1 "github.com/sazid/learngrpc/api/v1"
)

type LaptopStore interface {
	Save(*v1.Laptop) error
	Find(id string) (*v1.Laptop, error)
	Search(ctx context.Context, filter *v1.Filter, found func(laptop *v1.Laptop) error) error
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

	if _, ok := s.data[laptop.Id]; ok {
		return ErrAlreadyExists
	}

	other, err := deepCopy(laptop)
	if err != nil {
		return err
	}

	s.data[other.Id] = other

	return nil
}

func (s *InMemoryLaptopStore) Find(id string) (*v1.Laptop, error) {
	s.Lock()
	defer s.Unlock()
	if v, ok := s.data[id]; ok {
		return deepCopy(v)
	}
	return nil, ErrNotFound
}

func (s *InMemoryLaptopStore) Search(
	ctx context.Context,
	filter *v1.Filter,
	found func(laptop *v1.Laptop) error,
) error {
	s.RLock()
	defer s.RUnlock()

	for _, laptop := range s.data {
		log.Print("checking laptop id: ", laptop.GetId())
		if ctx.Err() == context.Canceled {
			log.Print("stream cancelled")
			return fmt.Errorf("context cancelled")
		}
		if ctx.Err() == context.DeadlineExceeded {
			log.Print("deadline exceeded")
			return fmt.Errorf("context cancelled")
		}
		if isQualified(filter, laptop) {
			other, err := deepCopy(laptop)
			if err != nil {
				return err
			}

			err = found(other)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isQualified(filter *v1.Filter, laptop *v1.Laptop) bool {
	if laptop.GetPriceUsd() > filter.MaxPriceUsd {
		return false
	}

	if laptop.Cpu.GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}

	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}

	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()) {
		return false
	}

	return true
}

func toBit(m *v1.Memory) uint64 {
	value := m.GetValue()
	switch m.GetUnit() {
	case v1.Memory_BIT:
		return value
	case v1.Memory_BYTE:
		return value << 3
	case v1.Memory_KILOBYTE:
		return value << 13
	case v1.Memory_MEGABYTE:
		return value << 23
	case v1.Memory_GIGABYTE:
		return value << 33
	case v1.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	}
}

func deepCopy(v *v1.Laptop) (*v1.Laptop, error) {
	other := &v1.Laptop{}
	err := copier.Copy(other, v)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return other, nil
}
