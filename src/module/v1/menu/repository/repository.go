package repository

import (
	"modular-monolithic/module/v1/menu/repository/postgresql"

	"git.motiolabs.com/library/motiolibs/mcarrier"
)

type MenuRepository struct {
	Carrier     *mcarrier.Carrier
	MenuPostgre postgresql.IMenuPostgre
}

func NewRepository(carrier *mcarrier.Carrier) MenuRepository {
	menuPostgre := postgresql.NewMenuPostgre(carrier)

	return MenuRepository{
		Carrier:     carrier,
		MenuPostgre: menuPostgre,
	}
}
