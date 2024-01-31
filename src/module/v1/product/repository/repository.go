package repository

import (
	"modular-monolithic/module/v1/product/repository/postgresql"

	"git.motiolabs.com/library/motiolibs/mcarrier"
)

type ProductRepository struct {
	Carrier        *mcarrier.Carrier
	ProductPostgre postgresql.IProductPostgre
}

func NewRepository(carrier *mcarrier.Carrier) ProductRepository {
	productPostgre := postgresql.NewProductPostgre(carrier)

	return ProductRepository{
		Carrier:        carrier,
		ProductPostgre: productPostgre,
	}
}
