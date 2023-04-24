package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsTopicCategoryModel = (*customCmsTopicCategoryModel)(nil)

type (
	// CmsTopicCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsTopicCategoryModel.
	CmsTopicCategoryModel interface {
		cmsTopicCategoryModel
	}

	customCmsTopicCategoryModel struct {
		*defaultCmsTopicCategoryModel
	}
)

// NewCmsTopicCategoryModel returns a model for the database table.
func NewCmsTopicCategoryModel(conn sqlx.SqlConn) CmsTopicCategoryModel {
	return &customCmsTopicCategoryModel{
		defaultCmsTopicCategoryModel: newCmsTopicCategoryModel(conn),
	}
}
