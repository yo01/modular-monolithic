package repository

import (
	authPostgre "modular-monolithic/module/v1/auth/repository/postgresql"

	"git.motiolabs.com/library/motiolibs/mcarrier"
)

type AuthRepository struct {
	Carrier     *mcarrier.Carrier
	AuthPostgre authPostgre.IAuthPostgre
}

func NewRepository(carrier *mcarrier.Carrier) AuthRepository {
	authPostgre := authPostgre.NewAuthPostgre(carrier)

	return AuthRepository{
		Carrier:     carrier,
		AuthPostgre: authPostgre,
	}
}
