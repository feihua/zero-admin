package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsMemberPrepaidCardLogModel = (*customUmsMemberPrepaidCardLogModel)(nil)

type (
	// UmsMemberPrepaidCardLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberPrepaidCardLogModel.
	UmsMemberPrepaidCardLogModel interface {
		umsMemberPrepaidCardLogModel
	}

	customUmsMemberPrepaidCardLogModel struct {
		*defaultUmsMemberPrepaidCardLogModel
	}
)

// NewUmsMemberPrepaidCardLogModel returns a model for the database table.
func NewUmsMemberPrepaidCardLogModel(conn sqlx.SqlConn) UmsMemberPrepaidCardLogModel {
	return &customUmsMemberPrepaidCardLogModel{
		defaultUmsMemberPrepaidCardLogModel: newUmsMemberPrepaidCardLogModel(conn),
	}
}
