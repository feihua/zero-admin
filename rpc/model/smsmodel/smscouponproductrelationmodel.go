package smsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SmsCouponProductRelationModel = (*customSmsCouponProductRelationModel)(nil)

type (
	// SmsCouponProductRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponProductRelationModel.
	SmsCouponProductRelationModel interface {
		smsCouponProductRelationModel
	}

	customSmsCouponProductRelationModel struct {
		*defaultSmsCouponProductRelationModel
	}
)

// NewSmsCouponProductRelationModel returns a model for the database table.
func NewSmsCouponProductRelationModel(conn sqlx.SqlConn) SmsCouponProductRelationModel {
	return &customSmsCouponProductRelationModel{
		defaultSmsCouponProductRelationModel: newSmsCouponProductRelationModel(conn),
	}
}
