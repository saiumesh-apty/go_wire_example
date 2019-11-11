package product

// This will be replaced by actual DB instance
type ProductRepository string

func ProvideProductRepository(input string) ProductRepository {
	return ProductRepository(input)
}

func (p *ProductRepository) GetProduct() Product {
	product := Product{
		Code:  "some_code_from_DB",
		Price: 1,
	}
	return product
}