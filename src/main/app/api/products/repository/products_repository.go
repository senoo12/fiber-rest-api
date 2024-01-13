package repository

import (
	domainProducts "Fiber/src/main/app/api/products/domain"
)

type ProductsRepository interface {
	CreateProducts(book domainProducts.Product) (domainProducts.Product, error)	
	GetAllProducts()([]domainProducts.Product, error)
	GetProductByID(productID int) (domainProducts.Product, error)
	UpdateProduct(productID int) (domainProducts.Product, error)
	DeleteProductByID(productID int) (domainProducts.Product, error)
}