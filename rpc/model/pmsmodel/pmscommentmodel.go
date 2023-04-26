package pmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmsCommentModel = (*customPmsCommentModel)(nil)

type (
	// PmsCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsCommentModel.
	PmsCommentModel interface {
		pmsCommentModel
	}

	customPmsCommentModel struct {
		*defaultPmsCommentModel
	}
)

// NewPmsCommentModel returns a model for the database table.
func NewPmsCommentModel(conn sqlx.SqlConn) PmsCommentModel {
	return &customPmsCommentModel{
		defaultPmsCommentModel: newPmsCommentModel(conn),
	}
}
