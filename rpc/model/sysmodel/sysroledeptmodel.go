package sysmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysRoleDeptModel = (*customSysRoleDeptModel)(nil)

type (
	// SysRoleDeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleDeptModel.
	SysRoleDeptModel interface {
		sysRoleDeptModel
	}

	customSysRoleDeptModel struct {
		*defaultSysRoleDeptModel
	}
)

// NewSysRoleDeptModel returns a model for the database table.
func NewSysRoleDeptModel(conn sqlx.SqlConn) SysRoleDeptModel {
	return &customSysRoleDeptModel{
		defaultSysRoleDeptModel: newSysRoleDeptModel(conn),
	}
}
