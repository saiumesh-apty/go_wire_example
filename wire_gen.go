// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"echo_wire/api/product"
)

// Injectors from wire.go:

func initProductAPI(input string) product.ProductAPI {
	productRepository := product.ProvideProductRepository(input)
	productService := product.ProvideProductService(productRepository)
	productAPI := product.ProvideProductAPI(productService)
	return productAPI
}