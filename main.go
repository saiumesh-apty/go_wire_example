package main

import "github.com/labstack/echo/v4"

func main() {
	fakeDBInstance := "AptyDB"
	productAPI := initProductAPI(fakeDBInstance)

	e := echo.New()
	e.GET("/", productAPI.GetProduct)
	e.Logger.Fatal(e.Start("localhost:1323"))

}