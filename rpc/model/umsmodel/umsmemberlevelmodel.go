package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsMemberLevelModel = (*customUmsMemberLevelModel)(nil)

type (
	// UmsMemberLevelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberLevelModel.
	UmsMemberLevelModel interface {
		umsMemberLevelModel
	}

	customUmsMemberLevelModel struct {
		*defaultUmsMemberLevelModel
	}
)

// NewUmsMemberLevelModel returns a model for the database table.
func NewUmsMemberLevelModel(conn sqlx.SqlConn) UmsMemberLevelModel {
	return &customUmsMemberLevelModel{
		defaultUmsMemberLevelModel: newUmsMemberLevelModel(conn),
	}
}
