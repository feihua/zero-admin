package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsMemberPrepaidCardRelationModel = (*customUmsMemberPrepaidCardRelationModel)(nil)

type (
	// UmsMemberPrepaidCardRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberPrepaidCardRelationModel.
	UmsMemberPrepaidCardRelationModel interface {
		umsMemberPrepaidCardRelationModel
	}

	customUmsMemberPrepaidCardRelationModel struct {
		*defaultUmsMemberPrepaidCardRelationModel
	}
)

// NewUmsMemberPrepaidCardRelationModel returns a model for the database table.
func NewUmsMemberPrepaidCardRelationModel(conn sqlx.SqlConn) UmsMemberPrepaidCardRelationModel {
	return &customUmsMemberPrepaidCardRelationModel{
		defaultUmsMemberPrepaidCardRelationModel: newUmsMemberPrepaidCardRelationModel(conn),
	}
}
