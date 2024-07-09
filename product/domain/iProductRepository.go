package domain

type IProductRepository interface {
	AddProduct(Product Product)
	GetProduct(id int) Product
}
