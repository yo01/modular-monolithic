package repository

import (
	"modular-monolithic/module/v1/cart/repository/postgresql"

	"git.motiolabs.com/library/motiolibs/mcarrier"
)

type CartRepository struct {
	Carrier         *mcarrier.Carrier
	CartPostgre     postgresql.ICartPostgre
	CartItemPostgre postgresql.ICartItemPostgre
}

func NewRepository(carrier *mcarrier.Carrier) CartRepository {
	cartPostgre := postgresql.NewCartPostgre(carrier)
	cartItemPostgre := postgresql.NewCartItemPostgre(carrier)

	return CartRepository{
		Carrier:         carrier,
		CartPostgre:     cartPostgre,
		CartItemPostgre: cartItemPostgre,
	}
}
