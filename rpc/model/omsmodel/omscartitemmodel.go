package omsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OmsCartItemModel = (*customOmsCartItemModel)(nil)

type (
	// OmsCartItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsCartItemModel.
	OmsCartItemModel interface {
		omsCartItemModel
	}

	customOmsCartItemModel struct {
		*defaultOmsCartItemModel
	}
)

// NewOmsCartItemModel returns a model for the database table.
func NewOmsCartItemModel(conn sqlx.SqlConn, c cache.CacheConf) OmsCartItemModel {
	return &customOmsCartItemModel{
		defaultOmsCartItemModel: newOmsCartItemModel(conn, c),
	}
}
