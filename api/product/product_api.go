package product

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductAPI struct {
	ProductService ProductService
}

func ProvideProductAPI(p ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

func(p *ProductAPI) GetProduct(c echo.Context) error {
	product := p.ProductService.FindProduct()
	return c.JSON(http.StatusOK, &product)
}