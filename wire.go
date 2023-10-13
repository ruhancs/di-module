//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/ruhancs/di-module/product"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	//cada vez que encontrar um productRepositoryInterface instacia o productRepository
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUsecase(db *sql.DB) *product.ProductUseCase{
	//criacao da injecao de dependencia
	wire.Build(
		//oque sera utilizado para criar a injecao de dependecia
		setRepositoryDependency,
		product.NewProductUseCase,
	)

	return &product.ProductUseCase{}
}

//anotations no topo, especie de tag para fazer a transformacao
//wire ira verificar as anotations