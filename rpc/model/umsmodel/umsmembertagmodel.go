package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsMemberTagModel = (*customUmsMemberTagModel)(nil)

type (
	// UmsMemberTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberTagModel.
	UmsMemberTagModel interface {
		umsMemberTagModel
	}

	customUmsMemberTagModel struct {
		*defaultUmsMemberTagModel
	}
)

// NewUmsMemberTagModel returns a model for the database table.
func NewUmsMemberTagModel(conn sqlx.SqlConn) UmsMemberTagModel {
	return &customUmsMemberTagModel{
		defaultUmsMemberTagModel: newUmsMemberTagModel(conn),
	}
}
