package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsSubjectCommentModel = (*customCmsSubjectCommentModel)(nil)

type (
	// CmsSubjectCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsSubjectCommentModel.
	CmsSubjectCommentModel interface {
		cmsSubjectCommentModel
	}

	customCmsSubjectCommentModel struct {
		*defaultCmsSubjectCommentModel
	}
)

// NewCmsSubjectCommentModel returns a model for the database table.
func NewCmsSubjectCommentModel(conn sqlx.SqlConn) CmsSubjectCommentModel {
	return &customCmsSubjectCommentModel{
		defaultCmsSubjectCommentModel: newCmsSubjectCommentModel(conn),
	}
}
