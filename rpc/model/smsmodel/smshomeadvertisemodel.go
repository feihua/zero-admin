package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sms/sms"
)

var _ SmsHomeAdvertiseModel = (*customSmsHomeAdvertiseModel)(nil)

type (
	// SmsHomeAdvertiseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeAdvertiseModel.
	SmsHomeAdvertiseModel interface {
		smsHomeAdvertiseModel
		Count(ctx context.Context, in *sms.HomeAdvertiseListReq) (int64, error)
		FindAll(ctx context.Context, in *sms.HomeAdvertiseListReq) (*[]SmsHomeAdvertise, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSmsHomeAdvertiseModel struct {
		*defaultSmsHomeAdvertiseModel
	}
)

// NewSmsHomeAdvertiseModel returns a model for the database table.
func NewSmsHomeAdvertiseModel(conn sqlx.SqlConn) SmsHomeAdvertiseModel {
	return &customSmsHomeAdvertiseModel{
		defaultSmsHomeAdvertiseModel: newSmsHomeAdvertiseModel(conn),
	}
}

func (m *customSmsHomeAdvertiseModel) FindAll(ctx context.Context, in *sms.HomeAdvertiseListReq) (*[]SmsHomeAdvertise, error) {

	where := "1=1"
	if in.Type != 2 {
		where = where + fmt.Sprintf(" AND type = %d", in.Type)
	}
	if in.Status != 2 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
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

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", smsHomeAdvertiseRows, m.table, where)
	var resp []SmsHomeAdvertise
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

func (m *customSmsHomeAdvertiseModel) Count(ctx context.Context, in *sms.HomeAdvertiseListReq) (int64, error) {
	where := "1=1"
	if in.Type != 2 {
		where = where + fmt.Sprintf(" AND type = %d", in.Type)
	}
	if in.Status != 2 {
		where = where + fmt.Sprintf(" AND status = %d", in.Status)
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

func (m *customSmsHomeAdvertiseModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
