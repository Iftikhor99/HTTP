package banners

import (
	"context"
	"errors"
	"sync"
)

// Service npenctasnset co6oi cepsuc no ynpasnenwo OaHHepamn.
type Service struct {
	mu sync.RWMutex

	items []*Banner
}

// NewService co3qaét cepsuc.
func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}
	//new comment
}

// Banner npenctasnaet codoi GaHHep.
type Banner struct {
	ID int64

	Title string

	Content string

	Button string

	Link string
}

// ByID Bo3BpawaeT OaHHep no upeHTHOuKaTopy.
func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}

	return nil, errors.New("item not found")
}

// All for
func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	panic("not implemented")
}

// Save for
func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {
	panic("not implemented")
}

// RemoveByID for
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	panic("not implemented")
}
