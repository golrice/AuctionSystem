package user

import (
	"context"
	"errors"
)

func (r *Repository) Create(ctx context.Context, model *Model) error {
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return errors.New("create user failed")
	}
	return nil
}

func (r *Repository) Get(ctx context.Context, model *Model) (*Model, error) {
	var result *Model
	if err := r.db.WithContext(ctx).Where(model).Take(&result).Error; err != nil {
		return nil, errors.New("get user failed")
	}

	return result, nil
}

func (r *Repository) Update(ctx context.Context, model *Model) error {
	if err := r.db.WithContext(ctx).Model(model).Updates(model).Error; err != nil {
		return errors.New("update user failed")
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, model *Model) error {
	if err := r.db.WithContext(ctx).Delete(model).Error; err != nil {
		return errors.New("delete user failed")
	}
	return nil
}
