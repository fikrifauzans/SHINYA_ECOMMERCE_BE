package persistence

import (
	"app/src/domain/product"
	"app/src/infrastructure/config"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository() product.ProductRepository {
	return &productRepository{
		db: config.DB,
	}
}

func (r *productRepository) FindAll(offset int, limit int) ([]product.Product, int64, error) {
	var products []product.Product
	var total int64
	r.db.Offset(offset).Limit(limit).Find(&products)
	r.db.Model(&product.Product{}).Count(&total)
	return products, total, nil
}

func (r *productRepository) FindByID(id uint) (product.Product, error) {
	var product product.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *productRepository) Create(product product.Product) (product.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *productRepository) Update(product product.Product) (product.Product, error) {
	if err := r.db.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *productRepository) Delete(id uint) error {
	if err := r.db.Delete(&product.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *productRepository) FindBySKU(email string) (product.Product, error) {
	var product product.Product
	if err := r.db.Where("SKU = ?", email).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
