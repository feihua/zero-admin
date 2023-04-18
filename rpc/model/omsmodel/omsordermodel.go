package omsmodel

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OmsOrderModel = (*customOmsOrderModel)(nil)

type (
	// OmsOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderModel.
	OmsOrderModel interface {
		omsOrderModel
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	customOmsOrderModel struct {
		*defaultOmsOrderModel
	}
)

// NewOmsOrderModel returns a model for the database table.
func NewOmsOrderModel(conn sqlx.SqlConn, c cache.CacheConf) OmsOrderModel {
	return &customOmsOrderModel{
		defaultOmsOrderModel: newOmsOrderModel(conn, c),
	}
}
func (m *customOmsOrderModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}
