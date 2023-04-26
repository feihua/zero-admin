package smsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SmsCouponProductCategoryRelationModel = (*customSmsCouponProductCategoryRelationModel)(nil)

type (
	// SmsCouponProductCategoryRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponProductCategoryRelationModel.
	SmsCouponProductCategoryRelationModel interface {
		smsCouponProductCategoryRelationModel
	}

	customSmsCouponProductCategoryRelationModel struct {
		*defaultSmsCouponProductCategoryRelationModel
	}
)

// NewSmsCouponProductCategoryRelationModel returns a model for the database table.
func NewSmsCouponProductCategoryRelationModel(conn sqlx.SqlConn) SmsCouponProductCategoryRelationModel {
	return &customSmsCouponProductCategoryRelationModel{
		defaultSmsCouponProductCategoryRelationModel: newSmsCouponProductCategoryRelationModel(conn),
	}
}
