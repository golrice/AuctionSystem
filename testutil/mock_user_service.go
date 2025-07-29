package testutil

import (
	"auctionsystem/internal/user"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MockUserRepository struct{}

func (m *MockUserRepository) Get(ctx context.Context, model *user.Model) (*user.Model, error) {
	if model.ID == 2 {
		return nil, errors.New("user not found")
	}
	return &user.Model{
		Model: gorm.Model{ID: 1},
		Name:  "golrice",
	}, nil
}

func (m *MockUserRepository) Update(ctx context.Context, model *user.Model) error { return nil }
func (m *MockUserRepository) Delete(ctx context.Context, model *user.Model) error { return nil }
func (m *MockUserRepository) Create(ctx context.Context, model *user.Model) error { return nil }

func NewMockUserRepository() user.UserRepository {
	return &MockUserRepository{}
}
