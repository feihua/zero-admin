package pmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmsCommentReplayModel = (*customPmsCommentReplayModel)(nil)

type (
	// PmsCommentReplayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsCommentReplayModel.
	PmsCommentReplayModel interface {
		pmsCommentReplayModel
	}

	customPmsCommentReplayModel struct {
		*defaultPmsCommentReplayModel
	}
)

// NewPmsCommentReplayModel returns a model for the database table.
func NewPmsCommentReplayModel(conn sqlx.SqlConn) PmsCommentReplayModel {
	return &customPmsCommentReplayModel{
		defaultPmsCommentReplayModel: newPmsCommentReplayModel(conn),
	}
}
