package cmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CmsMemberReportModel = (*customCmsMemberReportModel)(nil)

type (
	// CmsMemberReportModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCmsMemberReportModel.
	CmsMemberReportModel interface {
		cmsMemberReportModel
	}

	customCmsMemberReportModel struct {
		*defaultCmsMemberReportModel
	}
)

// NewCmsMemberReportModel returns a model for the database table.
func NewCmsMemberReportModel(conn sqlx.SqlConn) CmsMemberReportModel {
	return &customCmsMemberReportModel{
		defaultCmsMemberReportModel: newCmsMemberReportModel(conn),
	}
}
