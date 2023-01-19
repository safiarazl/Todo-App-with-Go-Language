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
	var categories []entity.Category
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("user_id = ?", id).Find(&categories).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []entity.Category{}, nil
		}
		return nil, ctx.Err()
	}
	return categories, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	err = r.db.WithContext(ctx).Model(&entity.Category{}).Create(&category).Error
	if err != nil {
		return 0, err
	}
	categoryId = category.ID
	// ragu
	return categoryId, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Create(&categories).Error
	if err != nil {
		return err
	}
	// wagu
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	var category entity.Category
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("id = ?", id).First(&category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Category{}, nil
		}
		return entity.Category{}, err
	}
	return category, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("id = ?", category.ID).Updates(&category)
	if err.Error != nil {
		return err.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Model(&entity.Category{}).Delete(&entity.Category{}, id)
	if err.Error != nil {
		return err.Error
	}
	return nil // TODO: replace this
}
