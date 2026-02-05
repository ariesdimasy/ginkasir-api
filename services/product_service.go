package services

import (
	"errors"
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

	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		CategoryID:  &req.CategoryID,
	}

	return ps.repo.CreateProduct(product)

}

func (ps *productService) UpdateProduct(id int64, req *models.UpdateProductRequest) error {

	if id <= 0 {
		return errors.New("invalid product id")
	}

	// Cek apakah product ada
	product, err := ps.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find product: %v", err)
	}
	if product == nil {
		return errors.New("product not found")
	}

	productRequest := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		CategoryID:  &req.CategoryID,
	}

	errUpdate := ps.repo.UpdateProduct(id, productRequest)

	if errUpdate != nil {
		return fmt.Errorf("failed to update product: %v ", errUpdate)
	}

	return productRequest
}

func (ps *productService) DeleteProduct(id int64) error {
	if id <= 0 {
		return errors.New("invalid product id")
	}

	// Cek apakah product ada
	product, err := ps.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find product: %v", err)
	}
	if product == nil {
		return errors.New("product not found")
	}

	errDelete := ps.repo.DeleteProduct(id)

	if errDelete != nil {
		return fmt.Errorf("failed to update product: %v ", errDelete)
	}

	return nil
}
