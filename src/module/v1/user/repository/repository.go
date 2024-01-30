package repository

import (
	userPostgre "modular-monolithic/module/v1/user/repository/postgresql"

	"git.motiolabs.com/library/motiolibs/mcarrier"
)

type UserRepository struct {
	Carrier     *mcarrier.Carrier
	UserPostgre userPostgre.IUserPostgre
}

func NewRepository(carrier *mcarrier.Carrier) UserRepository {

	userPostgre := userPostgre.NewUserPostgre(carrier)

	return UserRepository{
		Carrier:     carrier,
		UserPostgre: userPostgre,
	}
}
