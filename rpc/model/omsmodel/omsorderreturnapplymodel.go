package omsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/oms/oms"
)

var _ OmsOrderReturnApplyModel = (*customOmsOrderReturnApplyModel)(nil)

type (
	// OmsOrderReturnApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsOrderReturnApplyModel.
	OmsOrderReturnApplyModel interface {
		omsOrderReturnApplyModel
		Count(ctx context.Context, in *oms.OrderReturnApplyListReq) (int64, error)
		FindAll(ctx context.Context, in *oms.OrderReturnApplyListReq) (*[]OmsOrderReturnApply, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customOmsOrderReturnApplyModel struct {
		*defaultOmsOrderReturnApplyModel
	}
)

// NewOmsOrderReturnApplyModel returns a model for the database table.
func NewOmsOrderReturnApplyModel(conn sqlx.SqlConn) OmsOrderReturnApplyModel {
	return &customOmsOrderReturnApplyModel{
		defaultOmsOrderReturnApplyModel: newOmsOrderReturnApplyModel(conn),
	}
}

func (m *customOmsOrderReturnApplyModel) FindAll(ctx context.Context, in *oms.OrderReturnApplyListReq) (*[]OmsOrderReturnApply, error) {

	where := "1=1"
	if len(in.OrderSn) > 0 {
		where = where + fmt.Sprintf(" AND order_sn like '%%%s%%'", in.OrderSn)
	}
	if len(in.MemberUsername) > 0 {
		where = where + fmt.Sprintf(" AND member_username like '%%%s%%'", in.MemberUsername)
	}
	if in.Status != 4 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
	}
	if len(in.CreateTime) > 0 {
		where = where + fmt.Sprintf(" AND date_format(create_time,'%%Y-%%m-%%d') = '%s'", strings.Split(in.CreateTime, " ")[0])
	}
	if len(in.HandleTime) > 0 {
		where = where + fmt.Sprintf(" AND date_format(handle_time,'%%Y-%%m-%%d') = '%s'", strings.Split(in.HandleTime, " ")[0])
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", omsOrderReturnApplyRows, m.table, where)

	var resp []OmsOrderReturnApply
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

func (m *customOmsOrderReturnApplyModel) Count(ctx context.Context, in *oms.OrderReturnApplyListReq) (int64, error) {
	where := "1=1"
	if len(in.OrderSn) > 0 {
		where = where + fmt.Sprintf(" AND order_sn like '%%%s%%'", in.OrderSn)
	}
	if len(in.MemberUsername) > 0 {
		where = where + fmt.Sprintf(" AND member_username like '%%%s%%'", in.MemberUsername)
	}
	if in.Status != 4 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
	}
	if len(in.CreateTime) > 0 {
		where = where + fmt.Sprintf(" AND date_format(create_time,'%%Y-%%m-%%d') = '%s'", strings.Split(in.CreateTime, " ")[0])
	}
	if len(in.HandleTime) > 0 {
		where = where + fmt.Sprintf(" AND date_format(handle_time,'%%Y-%%m-%%d') = '%s'", strings.Split(in.HandleTime, " ")[0])
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

func (m *customOmsOrderReturnApplyModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
