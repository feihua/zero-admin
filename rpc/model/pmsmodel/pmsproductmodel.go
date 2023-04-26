package pmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmsProductModel = (*customPmsProductModel)(nil)

type (
	// PmsProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductModel.
	PmsProductModel interface {
		pmsProductModel
	}

	customPmsProductModel struct {
		*defaultPmsProductModel
	}
)

// NewPmsProductModel returns a model for the database table.
func NewPmsProductModel(conn sqlx.SqlConn) PmsProductModel {
	return &customPmsProductModel{
		defaultPmsProductModel: newPmsProductModel(conn),
	}
}
