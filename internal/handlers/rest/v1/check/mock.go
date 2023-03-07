package check

import (
	"context"
	"go.uber.org/zap"
)

type MockCheckUsecases struct {
}

func NewMockCheckUsecases() *MockCheckUsecases {
	return &MockCheckUsecases{}
}

func (m MockCheckUsecases) GetGeneratedCheckIDs(ctx context.Context) (GeneratedChecksResponse, error) {
	return GeneratedChecksResponse{
		IDs: []int{1, 3},
	}, nil
}

func (m MockCheckUsecases) SetChecksStatusPrinted(ctx context.Context, checkIDs []int) error {
	return nil
}

func (m MockCheckUsecases) GetCheckFilePath(ctx context.Context, checkID int) (string, error) {
	return "./testdata/0.pdf", nil
}

func mockHandler() *checkHandler {
	development, _ := zap.NewDevelopment()
	sugar := development.Sugar()
	return &checkHandler{
		log:      sugar,
		useCases: NewMockCheckUsecases(),
	}
}
