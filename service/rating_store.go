package service

import "sync"

type RatingStore interface {
	Add(laptopID string, score float64) (*Rating, error)
}

type Rating struct {
	Count uint32
	Sum   float64
}

type InMemoryRatingStore struct {
	sync.RWMutex
	rating map[string]*Rating
}

func NewInMemoryRatingStore() *InMemoryRatingStore {
	return &InMemoryRatingStore{
		rating: make(map[string]*Rating),
	}
}

func (s *InMemoryRatingStore) Add(laptopID string, score float64) (*Rating, error) {
	s.Lock()
	defer s.Unlock()

	rating, ok := s.rating[laptopID]
	if !ok {
		rating = &Rating{
			Count: 1,
			Sum:   score,
		}
	} else {
		rating.Count++
		rating.Sum += score
	}

	s.rating[laptopID] = rating

	return s.rating[laptopID], nil
}
