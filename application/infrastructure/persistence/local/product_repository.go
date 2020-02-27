package local

import (
	"errors"

	"github.com/maestre3d/bank-account/application/domain/model"
)

type ProductRepository struct {
	DB *DBPServer
}

var (
	productIDCount int64 = 0
)

func NewProductRepository(db *DBPServer) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (p *ProductRepository) Save(model *model.Product) error {
	p.DB.ProductDB = append(p.DB.ProductDB, model)
	return nil
}

func (p *ProductRepository) Update(model *model.Product) error {
	return nil
}

func (p *ProductRepository) Remove(ID int64) error {
	return nil
}

func (p *ProductRepository) FetchAll() ([]*model.Product, error) {
	products := p.DB.ProductDB
	if len(products) <= 0 {
		return nil, errors.New("products not found")
	}

	return products, nil
}

func (p *ProductRepository) FetchByID(ID int64) (*model.Product, error) {
	nID := p.DB.query().get(ID)
	selDB := DBPServer.query("select * from product where ID=?", nID)
	product := Product{}
	for selDB.Next() {
		var id float64
		var name string
		var price float64
		err = selDB.Scan(&id, &name, &price)
		if err != nil {
			panic(err.Error())
		}
		product.ID = id
		product.Name = name
		product.Price = price
		return product, nil
	}
}
