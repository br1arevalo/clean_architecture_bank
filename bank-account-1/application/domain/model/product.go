package model

import "errors"

type Product struct {
	ID    int64
	Name  string
	Price float32
}

func NewProduct(name string) *Product {
	return &Product{
		ID:    0,
		Name:  name,
		Price: 0.0,
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
