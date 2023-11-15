package service

import (
	"beleg-app/api/domain"
	"beleg-app/api/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetBelegById(t *testing.T) {
	mockRepo := new(mocks.BelegRepository)
	beleg := &domain.Beleg{Id: 1, Price: 100.0, Mwst: 7.8, Date: "31.12.2023", Shop: "Migros"}

	mockRepo.On("GetBelegById", mock.Anything, 1).Return(beleg, nil)

	service := NewBelegService(mockRepo)

	result, err := service.GetBelegById(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, beleg, result)

	mockRepo.AssertExpectations(t)
}

func TestGetAllBelege(t *testing.T) {
	mockRepo := new(mocks.BelegRepository)
	expectedBelege := []domain.Beleg{{Id: 1, Price: 100.0, Mwst: 7.8, Date: "31.12.2023", Shop: "Migros"}, {Id: 2, Price: 200.0, Mwst: 7.8, Date: "31.12.2023", Shop: "Coop"}}

	mockRepo.On("GetAllBelege", mock.Anything).Return(expectedBelege, nil)

	service := NewBelegService(mockRepo)

	result, err := service.GetAllBelege(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedBelege, result)

	mockRepo.AssertExpectations(t)
}
