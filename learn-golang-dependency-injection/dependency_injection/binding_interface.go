package dependency_injection

type SProduct struct {
	Name        string
	Description string
}

type IProduct interface {
	Add(product *SProduct) *SProduct
}

type ProductService struct {
	Product IProduct
}

type ProductImpl struct {
}

func (p *ProductImpl) Add(product *SProduct) *SProduct {
	return &SProduct{
		Name:        product.Name,
		Description: product.Description,
	}
}

func NewProductImpl() *ProductImpl {
	return &ProductImpl{}
}

func NewProductService(product IProduct) *ProductService {
	return &ProductService{
		Product: product,
	}
}
