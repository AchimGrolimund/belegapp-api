package mocks

import (
	"beleg-app/api/domain"
	"context"
	"github.com/stretchr/testify/mock"
)

type BelegRepository struct {
	mock.Mock
}

func (m *BelegRepository) CreateBeleg(ctx context.Context, beleg *domain.Beleg) error {
	args := m.Called(ctx, beleg)
	return args.Error(0)
}

func (m *BelegRepository) DeleteBelegById(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *BelegRepository) GetAllBelege(ctx context.Context) ([]domain.Beleg, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Beleg), args.Error(1)
}

func (m *BelegRepository) GetBelegById(ctx context.Context, id int) (*domain.Beleg, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Beleg), args.Error(1)
}
