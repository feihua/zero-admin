package sysmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysLogModel = (*customSysLogModel)(nil)

type (
	// SysLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysLogModel.
	SysLogModel interface {
		sysLogModel
	}

	customSysLogModel struct {
		*defaultSysLogModel
	}
)

// NewSysLogModel returns a model for the database table.
func NewSysLogModel(conn sqlx.SqlConn) SysLogModel {
	return &customSysLogModel{
		defaultSysLogModel: newSysLogModel(conn),
	}
}
