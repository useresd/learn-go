package account

import (
	"context"
	"fmt"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetByID(id int) (*Account, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Get(ctx context.Context, limit int, offset int, q string) ([]*Account, error) {
	return s.repo.Get(ctx, limit, offset, q)
}

func (s *Service) DeleteByID(ctx context.Context, id int) error {
	account, err := s.repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("couldn't find account by id %d", id)
	}
	return s.repo.Delete(ctx, account)
}
