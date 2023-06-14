package omsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ OmsOrderItemModel = (*customOmsOrderItemModel)(nil)

type (
	// OmsOrderItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderItemModel.
	OmsOrderItemModel interface {
		omsOrderItemModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64, OrderId int64) (*[]OmsOrderItem, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindProductListByOrderId(ctx context.Context, OrderId int64) (*[]OmsOrderItem, error)
	}

	customOmsOrderItemModel struct {
		*defaultOmsOrderItemModel
	}
)

// NewOmsOrderItemModel returns a model for the database table.
func NewOmsOrderItemModel(conn sqlx.SqlConn) OmsOrderItemModel {
	return &customOmsOrderItemModel{
		defaultOmsOrderItemModel: newOmsOrderItemModel(conn),
	}
}

func (m *customOmsOrderItemModel) FindAll(ctx context.Context, Current int64, PageSize int64, OrderId int64) (*[]OmsOrderItem, error) {

	query := fmt.Sprintf("select %s from %s where order_id = ? limit ?,?", omsOrderItemRows, m.table)
	var resp []OmsOrderItem
	err := m.conn.QueryRows(&resp, query, OrderId, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customOmsOrderItemModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

	var count int64
	err := m.conn.QueryRow(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customOmsOrderItemModel) DeleteByIds(ctx context.Context, ids []int64) error {
	// 拼接占位符 "?"
	placeholders := make([]string, len(ids))
	for i := range ids {
		placeholders[i] = "?"
	}

	// 构建删除语句
	query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", m.table, strings.Join(placeholders, ","))

	// 构建参数列表
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	// 执行删除语句
	_, err := m.conn.ExecCtx(ctx, query, args...)
	return err
}

func (m *customOmsOrderItemModel) FindProductListByOrderId(ctx context.Context, OrderId int64) (*[]OmsOrderItem, error) {

	query := fmt.Sprintf("select %s from %s where order_id = ?", omsOrderItemRows, m.table)
	var resp []OmsOrderItem
	err := m.conn.QueryRows(&resp, query, OrderId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
