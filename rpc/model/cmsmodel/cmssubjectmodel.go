package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsSubjectModel = (*customCmsSubjectModel)(nil)

type (
	// CmsSubjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsSubjectModel.
	CmsSubjectModel interface {
		cmsSubjectModel
	}

	customCmsSubjectModel struct {
		*defaultCmsSubjectModel
	}
)

// NewCmsSubjectModel returns a model for the database table.
func NewCmsSubjectModel(conn sqlx.SqlConn) CmsSubjectModel {
	return &customCmsSubjectModel{
		defaultCmsSubjectModel: newCmsSubjectModel(conn),
	}
}
