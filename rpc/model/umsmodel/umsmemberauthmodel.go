package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsMemberAuthModel = (*customUmsMemberAuthModel)(nil)

type (
	// UmsMemberAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberAuthModel.
	UmsMemberAuthModel interface {
		umsMemberAuthModel
	}

	customUmsMemberAuthModel struct {
		*defaultUmsMemberAuthModel
	}
)

// NewUmsMemberAuthModel returns a model for the database table.
func NewUmsMemberAuthModel(conn sqlx.SqlConn) UmsMemberAuthModel {
	return &customUmsMemberAuthModel{
		defaultUmsMemberAuthModel: newUmsMemberAuthModel(conn),
	}
}
