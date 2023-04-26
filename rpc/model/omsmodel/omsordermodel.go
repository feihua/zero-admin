package omsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OmsOrderModel = (*customOmsOrderModel)(nil)

type (
	// OmsOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderModel.
	OmsOrderModel interface {
		omsOrderModel
	}

	customOmsOrderModel struct {
		*defaultOmsOrderModel
	}
)

// NewOmsOrderModel returns a model for the database table.
func NewOmsOrderModel(conn sqlx.SqlConn) OmsOrderModel {
	return &customOmsOrderModel{
		defaultOmsOrderModel: newOmsOrderModel(conn),
	}
}
