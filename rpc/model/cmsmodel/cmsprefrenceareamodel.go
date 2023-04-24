package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsPrefrenceAreaModel = (*customCmsPrefrenceAreaModel)(nil)

type (
	// CmsPrefrenceAreaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsPrefrenceAreaModel.
	CmsPrefrenceAreaModel interface {
		cmsPrefrenceAreaModel
	}

	customCmsPrefrenceAreaModel struct {
		*defaultCmsPrefrenceAreaModel
	}
)

// NewCmsPrefrenceAreaModel returns a model for the database table.
func NewCmsPrefrenceAreaModel(conn sqlx.SqlConn) CmsPrefrenceAreaModel {
	return &customCmsPrefrenceAreaModel{
		defaultCmsPrefrenceAreaModel: newCmsPrefrenceAreaModel(conn),
	}
}
