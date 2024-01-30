package postgresql

import (
	"modular-monolithic/module/v1/role/repository/postgresql"

	"git.motiolabs.com/library/motiolibs/mcarrier"
)

type RoleRepository struct {
	Carrier     *mcarrier.Carrier
	RolePostgre postgresql.IRolePostgre
}

func NewRepository(carrier *mcarrier.Carrier) RoleRepository {
	rolePostgre := postgresql.NewRolePostgre(carrier)

	return RoleRepository{
		Carrier:     carrier,
		RolePostgre: rolePostgre,
	}
}
