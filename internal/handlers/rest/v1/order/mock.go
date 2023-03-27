package order

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"go.uber.org/zap"
)

type MockUsecases struct {
}

func (m MockUsecases) CreateChecks(ctx context.Context, order entity.Order) (ids []int, err error) {
	return []int{1, 2}, nil
}

func NewMockUsecases() *MockUsecases {
	return &MockUsecases{}
}

func mockHandler() *orderHandler {
	development, _ := zap.NewDevelopment()
	sugar := development.Sugar()
	return &orderHandler{
		log:      sugar,
		useCases: NewMockUsecases(),
	}
}
