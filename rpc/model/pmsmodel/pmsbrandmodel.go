package pmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmsBrandModel = (*customPmsBrandModel)(nil)

type (
	// PmsBrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsBrandModel.
	PmsBrandModel interface {
		pmsBrandModel
	}

	customPmsBrandModel struct {
		*defaultPmsBrandModel
	}
)

// NewPmsBrandModel returns a model for the database table.
func NewPmsBrandModel(conn sqlx.SqlConn) PmsBrandModel {
	return &customPmsBrandModel{
		defaultPmsBrandModel: newPmsBrandModel(conn),
	}
}
