package repositories

import (
	"ginkasir/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAllProducts(page int, limit int) ([]*models.Product, int64, error)
	FindByID(id int64) (*models.Product, error)
	FindByName(name string, page int, limit int) ([]*models.Product, int64, error)
	CreateProduct(req *models.CreateProductRequest) error
	UpdateProduct(id int64, req *models.UpdateProductRequest) error
	DeleteProduct(id int64) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (pr *productRepository) FindAllProducts(page int, limit int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64

	errCount := pr.db.Model(&models.Product{}).Count(&total).Error

	if errCount != nil {
		return nil, 0, errCount
	}

	offset := (page - 1) * limit
	errProduct := pr.db.
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&products)

	if errProduct != nil {
		return nil, 0, errProduct.Error
	}

	return products, total, nil

}

func (pr *productRepository) FindByID(id int64) (*models.Product, error) {
	var product *models.Product

	errProduct := pr.db.Where("id = ? ", id).First(&product).Error

	if errProduct != nil {
		return nil, errProduct
	}

	return product, nil
}

func (pr *productRepository) FindByName(name string, page int, limit int) ([]*models.Product, int64, error) {

	var products []*models.Product

	var total int64

	errCount := pr.db.Model(&models.Product{}).Count(&total).Error

	if errCount != nil {
		return nil, 0, errCount
	}

	offset := (page - 1) * limit

	errProducts := pr.db.Where("name like ?", "%"+name+"%").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&products)

	if errProducts != nil {
		return nil, 0, errProducts.Error
	}

	return products, total, nil

}

func (pr *productRepository) CreateProduct(req *models.CreateProductRequest) error {

	return pr.db.Create(&req).Error
}

func (pr *productRepository) UpdateProduct(id int64, req *models.UpdateProductRequest) error {
	return pr.db.Save(req).Where("id = ?", id).Error
}

func (pr *productRepository) DeleteProduct(id int64) error {
	return pr.db.Delete(id).Error
}
