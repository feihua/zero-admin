package omsmodel

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OmsOrderModel = (*customOmsOrderModel)(nil)

// var (
// 	cacheGozeroOmsOrderIdPrefix = "cache:gozero:omsOrder:id:"
// )

type (
	// OmsOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderModel.
	OmsOrderModel interface {
		omsOrderModel
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		FindListByMemberId(MemberId int64, Current int64, PageSize int64) (*[]OmsOrder, error)
		FindListByMemberIdAndStatus(MemberId int64, Status int64, Current int64, PageSize int64) (*[]OmsOrder, error)
		UpdateOrderStatus(ctx context.Context, status int64, id int64) error
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

// FindListByMemberId 根据用户id查询订单
func (m *customOmsOrderModel) FindListByMemberId(MemberId int64, Current int64, PageSize int64) (*[]OmsOrder, error) {

	query := fmt.Sprintf("select %s from %s where member_id = ? order by create_time desc limit ?,?", omsOrderRows, m.table)
	var resp []OmsOrder
	err := m.QueryRowsNoCache(&resp, query, MemberId, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customOmsOrderModel) FindListByMemberIdAndStatus(MemberId int64, Status int64, Current int64, PageSize int64) (*[]OmsOrder, error) {

	query := fmt.Sprintf("select %s from %s where member_id = ? and status = ? order by create_time desc limit ?,?", omsOrderRows, m.table)
	var resp []OmsOrder
	err := m.QueryRowsNoCache(&resp, query, MemberId, Status, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// UpdateOrderStatus 根据订单id更新订单状态
func (m *customOmsOrderModel) UpdateOrderStatus(ctx context.Context, status int64, id int64) error {
	// query := fmt.Sprintf("update %s set status = ? where id = ?", m.table)
	// _, err := m.ExecCtx(query, status, id)
	// return err

	gozeroOmsOrderIdKey := fmt.Sprintf("%s%v", cacheGozeroOmsOrderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set status = ? where id = ?", m.table)
		return conn.ExecCtx(ctx, query, status, id)
	}, gozeroOmsOrderIdKey)
	return err

}
