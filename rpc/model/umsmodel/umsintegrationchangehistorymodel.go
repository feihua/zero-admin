package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsIntegrationChangeHistoryModel = (*customUmsIntegrationChangeHistoryModel)(nil)

type (
	// UmsIntegrationChangeHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsIntegrationChangeHistoryModel.
	UmsIntegrationChangeHistoryModel interface {
		umsIntegrationChangeHistoryModel
	}

	customUmsIntegrationChangeHistoryModel struct {
		*defaultUmsIntegrationChangeHistoryModel
	}
)

// NewUmsIntegrationChangeHistoryModel returns a model for the database table.
func NewUmsIntegrationChangeHistoryModel(conn sqlx.SqlConn) UmsIntegrationChangeHistoryModel {
	return &customUmsIntegrationChangeHistoryModel{
		defaultUmsIntegrationChangeHistoryModel: newUmsIntegrationChangeHistoryModel(conn),
	}
}
