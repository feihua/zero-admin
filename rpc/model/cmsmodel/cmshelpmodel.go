package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsHelpModel = (*customCmsHelpModel)(nil)

type (
	// CmsHelpModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsHelpModel.
	CmsHelpModel interface {
		cmsHelpModel
	}

	customCmsHelpModel struct {
		*defaultCmsHelpModel
	}
)

// NewCmsHelpModel returns a model for the database table.
func NewCmsHelpModel(conn sqlx.SqlConn) CmsHelpModel {
	return &customCmsHelpModel{
		defaultCmsHelpModel: newCmsHelpModel(conn),
	}
}
