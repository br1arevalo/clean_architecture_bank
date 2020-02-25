package repository

import "github.com/maestre3d/bank-account/application/domain/model"

type IProductRepository interface {
	Save(model *model.Product) error
	Update(model *model.Product) error
	Remove(ID int64) error
	FetchAll() ([]*model.Product, error)
	FetchByID(ID int64) (*model.Product, error)
}
