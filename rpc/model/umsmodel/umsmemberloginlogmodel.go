package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsMemberLoginLogModel = (*customUmsMemberLoginLogModel)(nil)

type (
	// UmsMemberLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberLoginLogModel.
	UmsMemberLoginLogModel interface {
		umsMemberLoginLogModel
	}

	customUmsMemberLoginLogModel struct {
		*defaultUmsMemberLoginLogModel
	}
)

// NewUmsMemberLoginLogModel returns a model for the database table.
func NewUmsMemberLoginLogModel(conn sqlx.SqlConn) UmsMemberLoginLogModel {
	return &customUmsMemberLoginLogModel{
		defaultUmsMemberLoginLogModel: newUmsMemberLoginLogModel(conn),
	}
}
