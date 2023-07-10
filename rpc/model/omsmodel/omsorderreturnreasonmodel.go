package omsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/oms/omsclient"
)

var _ OmsOrderReturnReasonModel = (*customOmsOrderReturnReasonModel)(nil)

type (
	// OmsOrderReturnReasonModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderReturnReasonModel.
	OmsOrderReturnReasonModel interface {
		omsOrderReturnReasonModel
		Count(ctx context.Context, in *omsclient.OrderReturnReasonListReq) (int64, error)
		FindAll(ctx context.Context, in *omsclient.OrderReturnReasonListReq) (*[]OmsOrderReturnReason, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customOmsOrderReturnReasonModel struct {
		*defaultOmsOrderReturnReasonModel
	}
)

// NewOmsOrderReturnReasonModel returns a model for the database table.
func NewOmsOrderReturnReasonModel(conn sqlx.SqlConn) OmsOrderReturnReasonModel {
	return &customOmsOrderReturnReasonModel{
		defaultOmsOrderReturnReasonModel: newOmsOrderReturnReasonModel(conn),
	}
}

func (m *customOmsOrderReturnReasonModel) FindAll(ctx context.Context, in *omsclient.OrderReturnReasonListReq) (*[]OmsOrderReturnReason, error) {
	where := "1=1"
	if len(in.Name) > 0 {
		where = where + fmt.Sprintf(" AND name like '%%%s%%'", in.Name)
	}
	if in.Status != 2 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
	}
	query := fmt.Sprintf("select %s from %s where %s limit ?,?", omsOrderReturnReasonRows, m.table, where)
	var resp []OmsOrderReturnReason
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

func (m *customOmsOrderReturnReasonModel) Count(ctx context.Context, in *omsclient.OrderReturnReasonListReq) (int64, error) {
	where := "1=1"
	if len(in.Name) > 0 {
		where = where + fmt.Sprintf(" AND name like '%%%s%%'", in.Name)
	}
	if in.Status != 2 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
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

func (m *customOmsOrderReturnReasonModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
