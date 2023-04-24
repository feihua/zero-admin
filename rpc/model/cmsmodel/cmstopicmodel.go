package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsTopicModel = (*customCmsTopicModel)(nil)

type (
	// CmsTopicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsTopicModel.
	CmsTopicModel interface {
		cmsTopicModel
	}

	customCmsTopicModel struct {
		*defaultCmsTopicModel
	}
)

// NewCmsTopicModel returns a model for the database table.
func NewCmsTopicModel(conn sqlx.SqlConn) CmsTopicModel {
	return &customCmsTopicModel{
		defaultCmsTopicModel: newCmsTopicModel(conn),
	}
}
