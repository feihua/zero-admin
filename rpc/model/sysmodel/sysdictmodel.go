package sysmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SysDictModel = (*customSysDictModel)(nil)

type (
	// SysDictModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDictModel.
	SysDictModel interface {
		sysDictModel
	}

	customSysDictModel struct {
		*defaultSysDictModel
	}
)

// NewSysDictModel returns a model for the database table.
func NewSysDictModel(conn sqlx.SqlConn) SysDictModel {
	return &customSysDictModel{
		defaultSysDictModel: newSysDictModel(conn),
	}
}
