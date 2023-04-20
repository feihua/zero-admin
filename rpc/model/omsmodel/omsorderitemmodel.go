package omsmodel

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OmsOrderItemModel = (*customOmsOrderItemModel)(nil)

type (
	// OmsOrderItemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderItemModel.
	OmsOrderItemModel interface {
		omsOrderItemModel
		FindProductListByOrderId(OrderId int64) (*[]OmsOrderItem, error)
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

func (m *customOmsOrderItemModel) FindProductListByOrderId(OrderId int64) (*[]OmsOrderItem, error) {

	query := fmt.Sprintf("select %s from %s where order_id = ?", omsOrderItemRows, m.table)
	var resp []OmsOrderItem
	err := m.QueryRowsNoCache(&resp, query, OrderId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
