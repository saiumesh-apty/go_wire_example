//+build wireinject

package main

import (
	"echo_wire/api/product"
	"github.com/google/wire"
)

func initProductAPI(input string) product.ProductAPI {
	wire.Build(product.ProvideProductRepository, product.ProvideProductService, product.ProvideProductAPI)
	return product.ProductAPI{}
}