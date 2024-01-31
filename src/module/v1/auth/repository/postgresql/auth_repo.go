package postgresql

import "git.motiolabs.com/library/motiolibs/mcarrier"

type IAuthPostgre interface {
}

type authPostgre struct {
	Carrier *mcarrier.Carrier
}

func NewAuthPostgre(carrier *mcarrier.Carrier) authPostgre {
	return authPostgre{
		Carrier: carrier,
	}
}
