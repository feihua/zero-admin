package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsSubjectProductRelationModel = (*customCmsSubjectProductRelationModel)(nil)

type (
	// CmsSubjectProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsSubjectProductRelationModel.
	CmsSubjectProductRelationModel interface {
		cmsSubjectProductRelationModel
	}

	customCmsSubjectProductRelationModel struct {
		*defaultCmsSubjectProductRelationModel
	}
)

// NewCmsSubjectProductRelationModel returns a model for the database table.
func NewCmsSubjectProductRelationModel(conn sqlx.SqlConn) CmsSubjectProductRelationModel {
	return &customCmsSubjectProductRelationModel{
		defaultCmsSubjectProductRelationModel: newCmsSubjectProductRelationModel(conn),
	}
}
