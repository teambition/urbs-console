package dao

import (
	"github.com/teambition/urbs-console/src/service"
	"github.com/teambition/urbs-console/src/util"
)

var (
	daos *Daos
)

func init() {
	util.DigProvide(NewDaos)
}

// Daos ...
type Daos struct {
	OperationLog *OperationLog
	UrbsAcAcl    *UrbsAcAcl
	UrbsAcUser   *UrbsAcUser
}

// NewDaos ...
func NewDaos(sql *service.SQL) *Daos {
	daos := &Daos{
		OperationLog: &OperationLog{DB: sql.DB},
		UrbsAcAcl:    &UrbsAcAcl{DB: sql.DB},
		UrbsAcUser:   &UrbsAcUser{DB: sql.DB},
	}
	return daos
}