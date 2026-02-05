package repositories

import (
	"ginkasir/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(req *models.SearchProductRequest) ([]*models.Product, int64, error)
	FindByID(id int64) (*models.Product, error)
	FindByName(name string) (*models.Product, error)
	CreateProduct(req *models.Product) error
	UpdateProduct(id int64, req *models.Product) error
	DeleteProduct(id int64) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (pr *productRepository) FindAll(req *models.SearchProductRequest) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64

	errCount := pr.db.Model(&models.Product{}).Count(&total).Error

	if errCount != nil {
		return nil, 0, errCount
	}

	offset := (req.Page - 1) * req.Limit
	errData := pr.db.
		Preload("Category").
		Order("created_at DESC").
		Offset(offset).
		Limit(req.Limit)

	if req.Name != "" {
		errData = errData.Where("name ILIKE ? ", "%"+req.Name+"%")
	}

	errData = errData.Find(&products)

	if errData.Error != nil {
		return nil, 0, errData.Error
	}

	return products, total, nil

}

func (pr *productRepository) FindByID(id int64) (*models.Product, error) {
	var product *models.Product

	errProduct := pr.db.Preload("Category").Where("id = ? ", id).First(&product).Error

	if errProduct != nil {
		return nil, errProduct
	}

	return product, nil
}

func (pr *productRepository) FindByName(name string) (*models.Product, error) {

	var product *models.Product

	errProducts := pr.db.Where("name = ?", name).
		Find(&product)

	if errProducts != nil {
		return nil, errProducts.Error
	}

	return product, nil

}

func (pr *productRepository) CreateProduct(req *models.Product) error {

	return pr.db.Create(&req).Error
}

func (pr *productRepository) UpdateProduct(id int64, req *models.Product) error {
	return pr.db.Save(req).Where("id = ?", id).Error
}

func (pr *productRepository) DeleteProduct(id int64) error {
	return pr.db.Delete(id).Error
}
