package repositories

import (
	"ginkasir/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	FindByID(id int64) (*models.Category, error)
	FindAll(page int, limit int) ([]*models.Category, int64, error)
	FindByName(name string) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id int64) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (cr *categoryRepository) Create(category *models.Category) error {
	return cr.db.Create(category).Error
}

// Delete implements [CategoryRepository].
func (cr *categoryRepository) Delete(id int64) error {
	return cr.db.Delete(id).Error
}

// FindAll implements [CategoryRepository].
func (cr *categoryRepository) FindAll(page int, limit int) ([]*models.Category, int64, error) {

	var categories []*models.Category
	var total int64

	if errCount := cr.db.Model(&models.Category{}).Count(&total).Error; errCount != nil {
		return nil, 0, errCount
	}

	offset := (page - 1) * limit
	errData := cr.db.
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&categories).Error

	if errData != nil {
		return nil, 0, errData
	}

	return categories, total, nil

}

func (cr *categoryRepository) FindByName(name string) (*models.Category, error) {
	var category models.Category
	err := cr.db.Where("name = ?", name).First(&category).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

// FindByID implements [CategoryRepository].
func (cr *categoryRepository) FindByID(id int64) (*models.Category, error) {
	var category *models.Category

	err := cr.db.First(&category, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return category, nil

}

// Update implements [CategoryRepository].
func (cr *categoryRepository) Update(category *models.Category) error {
	return cr.db.Save(category).Error
}
