package usecase

import (
	"github.com/maestre3d/bank-account/application/domain/model"
	"github.com/maestre3d/bank-account/application/domain/repository"
)

type ProductUseCase struct {
	productRepository repository.IProductRepository
}

func NewProductUseCase(productRepository repository.IProductRepository) *ProductUseCase {
	return &ProductUseCase{productRepository}
}

func (p *ProductUseCase) CreateProduct(name string, price float64) error {
	product := model.NewProduct(name, price)
	if err := product.IsValid(); err != nil {
		return err
	}

	return p.productRepository.Save(product)
}

func (p *ProductUseCase) GetAllProducts() ([]*model.Product, error) {
	return p.productRepository.FetchAll()
}
