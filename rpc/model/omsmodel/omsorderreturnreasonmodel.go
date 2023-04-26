package omsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OmsOrderReturnReasonModel = (*customOmsOrderReturnReasonModel)(nil)

type (
	// OmsOrderReturnReasonModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderReturnReasonModel.
	OmsOrderReturnReasonModel interface {
		omsOrderReturnReasonModel
	}

	customOmsOrderReturnReasonModel struct {
		*defaultOmsOrderReturnReasonModel
	}
)

// NewOmsOrderReturnReasonModel returns a model for the database table.
func NewOmsOrderReturnReasonModel(conn sqlx.SqlConn) OmsOrderReturnReasonModel {
	return &customOmsOrderReturnReasonModel{
		defaultOmsOrderReturnReasonModel: newOmsOrderReturnReasonModel(conn),
	}
}
