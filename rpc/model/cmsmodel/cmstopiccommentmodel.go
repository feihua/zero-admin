package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsTopicCommentModel = (*customCmsTopicCommentModel)(nil)

type (
	// CmsTopicCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsTopicCommentModel.
	CmsTopicCommentModel interface {
		cmsTopicCommentModel
	}

	customCmsTopicCommentModel struct {
		*defaultCmsTopicCommentModel
	}
)

// NewCmsTopicCommentModel returns a model for the database table.
func NewCmsTopicCommentModel(conn sqlx.SqlConn) CmsTopicCommentModel {
	return &customCmsTopicCommentModel{
		defaultCmsTopicCommentModel: newCmsTopicCommentModel(conn),
	}
}
