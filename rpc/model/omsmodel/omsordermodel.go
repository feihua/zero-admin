package omsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/oms/oms"
)

var _ OmsOrderModel = (*customOmsOrderModel)(nil)

type (
	// OmsOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderModel.
	OmsOrderModel interface {
		omsOrderModel
		Count(ctx context.Context, in *oms.OrderListReq) (int64, error)
		FindAll(ctx context.Context, in *oms.OrderListReq) (*[]OmsOrder, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindListByMemberId(ctx context.Context, MemberId int64) (*[]OmsOrder, error)
	}

	customOmsOrderModel struct {
		*defaultOmsOrderModel
	}
)

// NewOmsOrderModel returns a model for the database table.
func NewOmsOrderModel(conn sqlx.SqlConn) OmsOrderModel {
	return &customOmsOrderModel{
		defaultOmsOrderModel: newOmsOrderModel(conn),
	}
}

func (m *customOmsOrderModel) FindAll(ctx context.Context, in *oms.OrderListReq) (*[]OmsOrder, error) {

	where := "1=1"
	if len(in.OrderSn) > 0 {
		where = where + fmt.Sprintf(" AND order_sn like '%%%s%%'", in.OrderSn)
	}
	if len(in.MemberUsername) > 0 {
		where = where + fmt.Sprintf(" AND member_username like '%%%s%%'", in.MemberUsername)
	}
	if in.PayType != 3 {
		where = where + fmt.Sprintf(" AND pay_type = %d", in.PayType)
	}
	if in.SourceType != 2 {
		where = where + fmt.Sprintf(" AND source_type = %d", in.SourceType)
	}
	if in.Status != 6 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
	}
	if in.OrderType != 2 {
		where = where + fmt.Sprintf(" AND order_type = %d", in.OrderType)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", omsOrderRows, m.table, where)
	var resp []OmsOrder
	err := m.conn.QueryRows(&resp, query, (in.Current-1)*in.PageSize, in.PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customOmsOrderModel) Count(ctx context.Context, in *oms.OrderListReq) (int64, error) {

	where := "1=1"
	if len(in.OrderSn) > 0 {
		where = where + fmt.Sprintf(" AND order_sn like '%%%s%%'", in.OrderSn)
	}
	if len(in.MemberUsername) > 0 {
		where = where + fmt.Sprintf(" AND member_username like '%%%s%%'", in.MemberUsername)
	}
	if in.PayType != 3 {
		where = where + fmt.Sprintf(" AND pay_type = %d", in.PayType)
	}
	if in.SourceType != 2 {
		where = where + fmt.Sprintf(" AND source_type = %d", in.SourceType)
	}
	if in.Status != 6 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
	}
	if in.OrderType != 2 {
		where = where + fmt.Sprintf(" AND order_type = %d", in.OrderType)
	}

	query := fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)

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

func (m *customOmsOrderModel) DeleteByIds(ctx context.Context, ids []int64) error {
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

// FindListByMemberId 根据用户id查询订单
func (m *customOmsOrderModel) FindListByMemberId(ctx context.Context, MemberId int64) (*[]OmsOrder, error) {

	query := fmt.Sprintf("select %s from %s where member_id = ? limit 1", omsOrderRows, m.table)
	var resp []OmsOrder
	err := m.conn.QueryRows(&resp, query, MemberId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
