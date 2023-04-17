package omsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OmsOrderOperateHistoryModel = (*customOmsOrderOperateHistoryModel)(nil)

type (
	// OmsOrderOperateHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderOperateHistoryModel.
	OmsOrderOperateHistoryModel interface {
		omsOrderOperateHistoryModel
	}

	customOmsOrderOperateHistoryModel struct {
		*defaultOmsOrderOperateHistoryModel
	}
)

// NewOmsOrderOperateHistoryModel returns a model for the database table.
func NewOmsOrderOperateHistoryModel(conn sqlx.SqlConn, c cache.CacheConf) OmsOrderOperateHistoryModel {
	return &customOmsOrderOperateHistoryModel{
		defaultOmsOrderOperateHistoryModel: newOmsOrderOperateHistoryModel(conn, c),
	}
}
