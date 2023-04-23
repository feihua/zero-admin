package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sms/sms"
)

var _ SmsCouponModel = (*customSmsCouponModel)(nil)

type (
	// SmsCouponModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponModel.
	SmsCouponModel interface {
		smsCouponModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64, in *sms.CouponListReq) (*[]SmsCoupon, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSmsCouponModel struct {
		*defaultSmsCouponModel
	}
)

// NewSmsCouponModel returns a model for the database table.
func NewSmsCouponModel(conn sqlx.SqlConn) SmsCouponModel {
	return &customSmsCouponModel{
		defaultSmsCouponModel: newSmsCouponModel(conn),
	}
}

func (m *customSmsCouponModel) FindAll(ctx context.Context, Current int64, PageSize int64, in *sms.CouponListReq) (*[]SmsCoupon, error) {

	where := "1=1"
	if in.Type != 4 {
		where = where + fmt.Sprintf(" AND type = %d", in.Type)
	}
	if in.Platform != 3 {
		where = where + fmt.Sprintf(" AND platform = %d", in.Platform)
	}
	if in.UseType != 3 {
		where = where + fmt.Sprintf(" AND use_type = %d", in.UseType)
	}
	if len(in.Name) > 0 {
		where = where + fmt.Sprintf(" AND name like '%%%s%%'", in.Name)
	}

	if len(in.StartTime) > 0 {
		where = where + fmt.Sprintf(" AND start_time >= '%s'", in.StartTime)
	}
	if len(in.EndTime) > 0 {
		where = where + fmt.Sprintf(" AND end_time <= '%s'", in.EndTime)
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", smsCouponRows, m.table, where)

	var resp []SmsCoupon
	err := m.conn.QueryRows(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSmsCouponModel) Count(ctx context.Context) (int64, error) {
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

func (m *customSmsCouponModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
