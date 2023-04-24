package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsHelpCategoryModel = (*customCmsHelpCategoryModel)(nil)

type (
	// CmsHelpCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsHelpCategoryModel.
	CmsHelpCategoryModel interface {
		cmsHelpCategoryModel
	}

	customCmsHelpCategoryModel struct {
		*defaultCmsHelpCategoryModel
	}
)

// NewCmsHelpCategoryModel returns a model for the database table.
func NewCmsHelpCategoryModel(conn sqlx.SqlConn) CmsHelpCategoryModel {
	return &customCmsHelpCategoryModel{
		defaultCmsHelpCategoryModel: newCmsHelpCategoryModel(conn),
	}
}
