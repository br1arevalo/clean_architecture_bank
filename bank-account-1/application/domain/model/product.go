package model

import "errors"

type Product struct {
	ID    int64
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    0,
		Name:  name,
		Price: price,
	}
}

func (p *Product) IsValid() error {
	if p.Price < 0 {
		return errors.New("invalid price")
	} else if p.Name == "" {
		return errors.New("invalid name")
	}

	return nil
}
