package product

type ProductService struct {
	ProductRepository ProductRepository
}

func ProvideProductService(p ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) FindProduct() Product {
	return p.ProductRepository.GetProduct()
}