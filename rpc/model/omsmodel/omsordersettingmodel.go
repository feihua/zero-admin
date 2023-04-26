package omsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OmsOrderSettingModel = (*customOmsOrderSettingModel)(nil)

type (
	// OmsOrderSettingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderSettingModel.
	OmsOrderSettingModel interface {
		omsOrderSettingModel
	}

	customOmsOrderSettingModel struct {
		*defaultOmsOrderSettingModel
	}
)

// NewOmsOrderSettingModel returns a model for the database table.
func NewOmsOrderSettingModel(conn sqlx.SqlConn) OmsOrderSettingModel {
	return &customOmsOrderSettingModel{
		defaultOmsOrderSettingModel: newOmsOrderSettingModel(conn),
	}
}
