package local

import "github.com/maestre3d/bank-account/application/domain/model"

type DBServer struct {
	UserDB []*model.User
}

type DBPServer struct {
	ProductDB []*model.Product
}
