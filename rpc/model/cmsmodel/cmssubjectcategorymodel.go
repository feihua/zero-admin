package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsSubjectCategoryModel = (*customCmsSubjectCategoryModel)(nil)

type (
	// CmsSubjectCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsSubjectCategoryModel.
	CmsSubjectCategoryModel interface {
		cmsSubjectCategoryModel
	}

	customCmsSubjectCategoryModel struct {
		*defaultCmsSubjectCategoryModel
	}
)

// NewCmsSubjectCategoryModel returns a model for the database table.
func NewCmsSubjectCategoryModel(conn sqlx.SqlConn) CmsSubjectCategoryModel {
	return &customCmsSubjectCategoryModel{
		defaultCmsSubjectCategoryModel: newCmsSubjectCategoryModel(conn),
	}
}
