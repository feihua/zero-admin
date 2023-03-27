package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsMemberModel = (*customUmsMemberModel)(nil)

type (
	// UmsMemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberModel.
	UmsMemberModel interface {
		umsMemberModel
	}

	customUmsMemberModel struct {
		*defaultUmsMemberModel
	}
)

// NewUmsMemberModel returns a model for the database table.
func NewUmsMemberModel(conn sqlx.SqlConn) UmsMemberModel {
	return &customUmsMemberModel{
		defaultUmsMemberModel: newUmsMemberModel(conn),
	}
}
