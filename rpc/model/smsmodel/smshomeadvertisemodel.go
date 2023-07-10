package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"zero-admin/rpc/sms/smsclient"
)

var _ SmsHomeAdvertiseModel = (*customSmsHomeAdvertiseModel)(nil)

type (
	// SmsHomeAdvertiseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsHomeAdvertiseModel.
	SmsHomeAdvertiseModel interface {
		smsHomeAdvertiseModel
		Count(ctx context.Context, in *smsclient.HomeAdvertiseListReq) (int64, error)
		FindAll(ctx context.Context, in *smsclient.HomeAdvertiseListReq) (*[]SmsHomeAdvertise, error)
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

func (m *customSmsHomeAdvertiseModel) FindAll(ctx context.Context, in *smsclient.HomeAdvertiseListReq) (*[]SmsHomeAdvertise, error) {

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

func (m *customSmsHomeAdvertiseModel) Count(ctx context.Context, in *smsclient.HomeAdvertiseListReq) (int64, error) {
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
