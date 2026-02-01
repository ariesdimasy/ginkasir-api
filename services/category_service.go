package services

import (
	"errors"
	"fmt"
	"ginkasir/models"
	"ginkasir/repositories"
)

type CategoryService interface {
	CreateCategory(req *models.CreateCategoryRequest) (*models.Category, error)
	GetCategoryByID(id int64) (*models.Category, error)
	GetAllCategories(page, limit int) ([]*models.Category, int64, error)
	UpdateCategory(id int64, req *models.UpdateCategoryRequest) (*models.Category, error)
	DeleteCategory(id int64) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (cs *categoryService) CreateCategory(req *models.CreateCategoryRequest) (*models.Category, error) {

	existing, err := cs.repo.FindByName(req.Name)

	if err != nil {
		return nil, fmt.Errorf("Failed to check duplicate category : %v", err)
	}

	if existing != nil {
		return nil, errors.New("Category name already exists")
	}

	category := &models.Category{
		Name: req.Name,
	}

	// if errValidate := category.Validate(); errValidate != nil {
	// 	return nil, err
	// }

	if errRepository := cs.repo.Create(category); errRepository != nil {
		return nil, fmt.Errorf("failed to create category: %v ", err)
	}

	return category, nil

}

func (cs *categoryService) GetCategoryByID(id int64) (*models.Category, error) {
	if id <= 0 {
		return nil, errors.New("Invalid Category ID")
	}

	category, err := cs.repo.FindByID(id)

	if err != nil {
		return nil, fmt.Errorf("failed to get Category : %v", err)
	}
	if category == nil {
		return nil, errors.New("Category not found")
	}

	return category, nil

}

func (cs *categoryService) GetAllCategories(page, limit int) ([]*models.Category, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	categories, total, err := cs.repo.FindAll(page, limit)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get categories: %v", err)
	}

	return categories, total, nil

}

func (cs *categoryService) UpdateCategory(id int64, req *models.UpdateCategoryRequest) (*models.Category, error) {
	category, err := cs.repo.FindByID(id)

	if err != nil {
		return nil, fmt.Errorf("failed to get Category : %v", err)
	}
	if category == nil {
		return nil, errors.New("Category not found")
	}

	if category.Name != req.Name {
		existing, err := cs.repo.FindByName(req.Name)
		if err != nil {
			return nil, fmt.Errorf("Failed to check duplicate category : %v", err)
		}

		if existing != nil && int64(existing.ID) != id {
			return nil, errors.New("Category name already exists")
		}
	}

	category.Name = req.Name

	if err := cs.repo.Update(category); err != nil {
		return nil, fmt.Errorf("failed to update category : %v ", err)
	}

	return category, nil
}

func (cs *categoryService) DeleteCategory(id int64) error {
	if id <= 0 {
		return errors.New("invalid category id")
	}

	// Cek apakah category ada
	category, err := cs.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find category: %v", err)
	}
	if category == nil {
		return errors.New("category not found")
	}

	// Business rule: Cek apakah category masih digunakan
	// Contoh: if category.HasProducts() { return error }

	// Hapus category
	if err := cs.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete category: %v", err)
	}

	return nil
}
