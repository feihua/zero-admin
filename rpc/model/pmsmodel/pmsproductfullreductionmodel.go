package pmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmsProductFullReductionModel = (*customPmsProductFullReductionModel)(nil)

type (
	// PmsProductFullReductionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductFullReductionModel.
	PmsProductFullReductionModel interface {
		pmsProductFullReductionModel
	}

	customPmsProductFullReductionModel struct {
		*defaultPmsProductFullReductionModel
	}
)

// NewPmsProductFullReductionModel returns a model for the database table.
func NewPmsProductFullReductionModel(conn sqlx.SqlConn) PmsProductFullReductionModel {
	return &customPmsProductFullReductionModel{
		defaultPmsProductFullReductionModel: newPmsProductFullReductionModel(conn),
	}
}
