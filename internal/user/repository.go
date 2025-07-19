package user

import (
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, model *UserModel) error {
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *userRepository) Get(ctx context.Context, model *UserModel) (*UserModel, error) {
	var result *UserModel
	err := r.db.WithContext(ctx).Where(model).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepository) Update(ctx context.Context, model *UserModel) error {
	return r.db.WithContext(ctx).Model(model).Updates(model).Error
}

func (r *userRepository) Delete(ctx context.Context, model *UserModel) error {
	return r.db.WithContext(ctx).Delete(model).Error
}
