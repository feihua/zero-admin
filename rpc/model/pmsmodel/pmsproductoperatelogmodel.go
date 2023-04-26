package pmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmsProductOperateLogModel = (*customPmsProductOperateLogModel)(nil)

type (
	// PmsProductOperateLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductOperateLogModel.
	PmsProductOperateLogModel interface {
		pmsProductOperateLogModel
	}

	customPmsProductOperateLogModel struct {
		*defaultPmsProductOperateLogModel
	}
)

// NewPmsProductOperateLogModel returns a model for the database table.
func NewPmsProductOperateLogModel(conn sqlx.SqlConn) PmsProductOperateLogModel {
	return &customPmsProductOperateLogModel{
		defaultPmsProductOperateLogModel: newPmsProductOperateLogModel(conn),
	}
}
