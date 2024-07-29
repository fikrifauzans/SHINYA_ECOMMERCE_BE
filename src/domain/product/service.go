package product

type ProductService interface {
	GetProducts(page int, perPage int) ([]Product, int64, error)
	GetProductByID(id uint) (Product, error)
	CreateProduct(product Product) (Product, error)
	UpdateProduct(product Product) (Product, error)
	DeleteProduct(id uint) error
	FindProductBySKU(email string) (Product, error)
}
