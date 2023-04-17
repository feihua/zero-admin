package omsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OmsOrderReturnApplyModel = (*customOmsOrderReturnApplyModel)(nil)

type (
	// OmsOrderReturnApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderReturnApplyModel.
	OmsOrderReturnApplyModel interface {
		omsOrderReturnApplyModel
	}

	customOmsOrderReturnApplyModel struct {
		*defaultOmsOrderReturnApplyModel
	}
)

// NewOmsOrderReturnApplyModel returns a model for the database table.
func NewOmsOrderReturnApplyModel(conn sqlx.SqlConn, c cache.CacheConf) OmsOrderReturnApplyModel {
	return &customOmsOrderReturnApplyModel{
		defaultOmsOrderReturnApplyModel: newOmsOrderReturnApplyModel(conn, c),
	}
}
