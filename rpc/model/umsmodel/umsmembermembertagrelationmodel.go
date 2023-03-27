package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsMemberMemberTagRelationModel = (*customUmsMemberMemberTagRelationModel)(nil)

type (
	// UmsMemberMemberTagRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberMemberTagRelationModel.
	UmsMemberMemberTagRelationModel interface {
		umsMemberMemberTagRelationModel
	}

	customUmsMemberMemberTagRelationModel struct {
		*defaultUmsMemberMemberTagRelationModel
	}
)

// NewUmsMemberMemberTagRelationModel returns a model for the database table.
func NewUmsMemberMemberTagRelationModel(conn sqlx.SqlConn) UmsMemberMemberTagRelationModel {
	return &customUmsMemberMemberTagRelationModel{
		defaultUmsMemberMemberTagRelationModel: newUmsMemberMemberTagRelationModel(conn),
	}
}
