package pmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmsProductCollectModel = (*customPmsProductCollectModel)(nil)

type (
	// PmsProductCollectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductCollectModel.
	PmsProductCollectModel interface {
		pmsProductCollectModel
	}

	customPmsProductCollectModel struct {
		*defaultPmsProductCollectModel
	}
)

// NewPmsProductCollectModel returns a model for the database table.
func NewPmsProductCollectModel(conn sqlx.SqlConn) PmsProductCollectModel {
	return &customPmsProductCollectModel{
		defaultPmsProductCollectModel: newPmsProductCollectModel(conn),
	}
}
