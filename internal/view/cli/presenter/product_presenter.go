package presenter

import (
	"github.com/maestre3d/bank-account/application/domain/model"
	"github.com/maestre3d/bank-account/application/infrastructure/persistence/local"
	"github.com/maestre3d/bank-account/application/usecase"
)

type ProductPresenter struct {
	useCase *usecase.ProductUseCase
}

func NewProductPresenter() *ProductPresenter {
	products := make([]*model.Product, 0)
	db := local.DBPServer{ProductDB: products}
	useCase := usecase.NewProductUseCase(local.NewProductRepository(&db))

	return &ProductPresenter{useCase: useCase}
}

func (p *ProductPresenter) CreateProduct(name string, price float64) error {
	return p.useCase.CreateProduct(name, price)
}

func (p *ProductPresenter) GetAll() []*model.Product {
	products, err := p.useCase.GetAllProducts()
	if err != nil || len(products) <= 0 {
		return nil
	}

	return products
}

func (p *ProductPresenter) GetByID(ID float64) *model.Product {
	product :=
}
