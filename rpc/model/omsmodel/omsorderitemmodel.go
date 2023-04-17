package omsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OmsOrderItemModel = (*customOmsOrderItemModel)(nil)

type (
	// OmsOrderItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderItemModel.
	OmsOrderItemModel interface {
		omsOrderItemModel
	}

	customOmsOrderItemModel struct {
		*defaultOmsOrderItemModel
	}
)

// NewOmsOrderItemModel returns a model for the database table.
func NewOmsOrderItemModel(conn sqlx.SqlConn, c cache.CacheConf) OmsOrderItemModel {
	return &customOmsOrderItemModel{
		defaultOmsOrderItemModel: newOmsOrderItemModel(conn, c),
	}
}
