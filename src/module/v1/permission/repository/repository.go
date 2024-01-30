package repository

import (
	"modular-monolithic/module/v1/permission/repository/postgresql"

	"git.motiolabs.com/library/motiolibs/mcarrier"
)

type PermissionRepository struct {
	Carrier           *mcarrier.Carrier
	PermissionPostgre postgresql.IPermissionPostgre
}

func NewRepository(carrier *mcarrier.Carrier) PermissionRepository {
	permissionPostgre := postgresql.NewPermissionPostgre(carrier)

	return PermissionRepository{
		Carrier:           carrier,
		PermissionPostgre: permissionPostgre,
	}
}
