package umsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UmsGrowthChangeHistoryModel = (*customUmsGrowthChangeHistoryModel)(nil)

type (
	// UmsGrowthChangeHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsGrowthChangeHistoryModel.
	UmsGrowthChangeHistoryModel interface {
		umsGrowthChangeHistoryModel
	}

	customUmsGrowthChangeHistoryModel struct {
		*defaultUmsGrowthChangeHistoryModel
	}
)

// NewUmsGrowthChangeHistoryModel returns a model for the database table.
func NewUmsGrowthChangeHistoryModel(conn sqlx.SqlConn) UmsGrowthChangeHistoryModel {
	return &customUmsGrowthChangeHistoryModel{
		defaultUmsGrowthChangeHistoryModel: newUmsGrowthChangeHistoryModel(conn),
	}
}
