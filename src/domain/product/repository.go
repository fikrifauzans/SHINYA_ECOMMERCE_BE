package product

type ProductRepository interface {
	FindAll(offset int, limit int) ([]Product, int64, error)
	FindByID(id uint) (Product, error)
	Create(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(id uint) error
	FindBySKU(sku string) (Product, error)
}
