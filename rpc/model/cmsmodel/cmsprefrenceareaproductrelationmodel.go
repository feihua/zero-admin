package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsPrefrenceAreaProductRelationModel = (*customCmsPrefrenceAreaProductRelationModel)(nil)

type (
	// CmsPrefrenceAreaProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsPrefrenceAreaProductRelationModel.
	CmsPrefrenceAreaProductRelationModel interface {
		cmsPrefrenceAreaProductRelationModel
	}

	customCmsPrefrenceAreaProductRelationModel struct {
		*defaultCmsPrefrenceAreaProductRelationModel
	}
)

// NewCmsPrefrenceAreaProductRelationModel returns a model for the database table.
func NewCmsPrefrenceAreaProductRelationModel(conn sqlx.SqlConn) CmsPrefrenceAreaProductRelationModel {
	return &customCmsPrefrenceAreaProductRelationModel{
		defaultCmsPrefrenceAreaProductRelationModel: newCmsPrefrenceAreaProductRelationModel(conn),
	}
}
