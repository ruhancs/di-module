package product

import "database/sql"

type ProductRepositoryInterface interface {
	GetProduct(id int) Product
}

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (repo *ProductRepository) GetProduct(id int) Product {
	return Product{
		ID: id,
		Name: "P1",
	}
}