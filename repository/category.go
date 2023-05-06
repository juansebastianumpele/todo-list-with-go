package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	categoryData := []entity.Category{}
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("user_id = ?", id).Find(&categoryData)
	if err.Error != nil {
		return []entity.Category{}, err.Error
	}
	return categoryData, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	res := r.db.WithContext(ctx).Create(&category)
	if res.Error != nil {
		return 0, res.Error
	}
	return category.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	res := r.db.WithContext(ctx).Create(&categories)
	if res.Error != nil {
		return res.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	categoryData := entity.Category{}
	res := r.db.WithContext(ctx).Model(&entity.Category{}).Where("id = ?", id).Find(&categoryData)
	if res.Error != nil {
		return entity.Category{}, res.Error
	}
	return categoryData, nil // TODO: replace this
	// return entity.Category{}, nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	res := r.db.WithContext(ctx).Model(&entity.Category{}).Where("id = ?", category.ID).Updates(&category)
	if res.Error != nil {
		return res.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	res := r.db.WithContext(ctx).Model(&entity.Category{}).Where("id = ?", id).Delete(&entity.Category{})
	if res.Error != nil {
		return res.Error
	}
	return nil // TODO: replace this
}
