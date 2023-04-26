package smsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SmsCouponHistoryModel = (*customSmsCouponHistoryModel)(nil)

type (
	// SmsCouponHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponHistoryModel.
	SmsCouponHistoryModel interface {
		smsCouponHistoryModel
	}

	customSmsCouponHistoryModel struct {
		*defaultSmsCouponHistoryModel
	}
)

// NewSmsCouponHistoryModel returns a model for the database table.
func NewSmsCouponHistoryModel(conn sqlx.SqlConn) SmsCouponHistoryModel {
	return &customSmsCouponHistoryModel{
		defaultSmsCouponHistoryModel: newSmsCouponHistoryModel(conn),
	}
}
