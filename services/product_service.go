package services

import (
	"fmt"
	"ginkasir/models"
	"ginkasir/repositories"
)

type ProductService interface {
	CreateProduct(req *models.CreateProductRequest) error
	UpdateProduct(id int64, req *models.UpdateProductRequest) error
	DeleteProduct(id int64) error
	GetAllProducts(query *models.SearchProductRequest) ([]*models.Product, int64, error)
	GetProductByID(id int64) (*models.Product, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *productService {
	return &productService{repo: repo}
}

func (ps *productService) GetAllProducts(req *models.SearchProductRequest) ([]*models.Product, int64, error) {

	if req.Page < 1 {
		req.Page = 1
	}

	if req.Limit < 1 || req.Limit > 100 {
		req.Limit = 10
	}

	var products []*models.Product
	var total int64
	var err error

	products, total, err = ps.repo.FindAll(req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get categories: %v", err)
	}

	return products, total, nil

}

func (ps *productService) GetProductByID(id int64) (*models.Product, error) {

	var product *models.Product
	var err error

	product, err = ps.repo.FindByID(id)

	if err != nil {
		return nil, fmt.Errorf("product doesnt exists : %v", err)
	}

	return product, nil
}

func (ps *productService) CreateProduct(req *models.CreateProductRequest) error {

	existing, err := ps.repo.FindByName(req.Name)

	if err != nil {
		return fmt.Errorf("Failed to check duplicate product name")
	}

	if existing != nil {
		return fmt.Errorf("product name already exists ")
	}

	return ps.repo.CreateProduct(req)

}

func (ps *productService) UpdateProduct(id int64, req *models.UpdateProductRequest) error {
	return ps.repo.UpdateProduct(id, req)
}

func (ps *productService) DeleteProduct(id int64) error {
	return ps.repo.DeleteProduct(id)
}
