package umsmodel

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsMemberModel = (*customUmsMemberModel)(nil)

type (
	// UmsMemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberModel.
	UmsMemberModel interface {
		umsMemberModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
	}

	customUmsMemberModel struct {
		*defaultUmsMemberModel
	}
)

// NewUmsMemberModel returns a model for the database table.
func NewUmsMemberModel(conn sqlx.SqlConn, c cache.CacheConf) UmsMemberModel {
	return &customUmsMemberModel{
		defaultUmsMemberModel: newUmsMemberModel(conn, c),
	}
}

func (m *defaultUmsMemberModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}
