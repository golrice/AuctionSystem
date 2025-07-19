package auction

import (
	"context"
	"errors"
)

func (r *Repository) Create(ctx context.Context, model *Model) error {
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return errors.New("create auction failed")
	}
	return nil
}

func (r *Repository) Get(ctx context.Context, model *Model) (*Model, error) {
	var result *Model
	if err := r.db.WithContext(ctx).Where(model).First(&result).Error; err != nil {
		return nil, errors.New("get auction failed")
	}
	return result, nil
}

func (r *Repository) Update(ctx context.Context, model *Model) error {
	if err := r.db.WithContext(ctx).Model(model).Updates(model).Error; err != nil {
		return errors.New("update auction failed")
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, model *Model) error {
	if err := r.db.WithContext(ctx).Delete(model).Error; err != nil {
		return errors.New("delete auction failed")
	}
	return nil
}

func (r *Repository) List(ctx context.Context, model *Model) ([]*Model, error) {
	var result []*Model
	if err := r.db.WithContext(ctx).Where(model).Find(&result).Error; err != nil {
		return nil, errors.New("list auction failed")
	}
	return result, nil
}
