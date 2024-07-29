package usecase

import (
	"app/src/domain/product"
)

type productUsecase struct {
	productRepository product.ProductRepository
}

func NewProductUsecase(r product.ProductRepository) product.ProductService {
	return &productUsecase{
		productRepository: r,
	}
}

func (u *productUsecase) GetProducts(page int, perPage int) ([]product.Product, int64, error) {
	offset := (page - 1) * perPage
	return u.productRepository.FindAll(offset, perPage)
}

func (u *productUsecase) GetProductByID(id uint) (product.Product, error) {
	return u.productRepository.FindByID(id)
}

func (u *productUsecase) CreateProduct(user product.Product) (product.Product, error) {
	return u.productRepository.Create(user)
}

func (u *productUsecase) UpdateProduct(user product.Product) (product.Product, error) {
	return u.productRepository.Update(user)
}

func (u *productUsecase) DeleteProduct(id uint) error {
	return u.productRepository.Delete(id)
}

func (u *productUsecase) FindProductBySKU(sku string) (product.Product, error) {
	return u.productRepository.FindBySKU(sku)
}
