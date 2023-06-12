package smsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ SmsCouponHistoryModel = (*customSmsCouponHistoryModel)(nil)

type (
	// SmsCouponHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsCouponHistoryModel.
	SmsCouponHistoryModel interface {
		smsCouponHistoryModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SmsCouponHistory, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSmsCouponHistoryModel struct {
		*defaultSmsCouponHistoryModel
	}
)

// NewSmsCouponHistoryModel returns a model for the database table.
func NewSmsCouponHistoryModel(conn sqlx.SqlConn) SmsCouponHistoryModel {
	return &customSmsCouponHistoryModel{
		defaultSmsCouponHistoryModel: newSmsCouponHistoryModel(conn),
	}
}

func (m *customSmsCouponHistoryModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SmsCouponHistory, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", smsCouponHistoryRows, m.table)
	var resp []SmsCouponHistory
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

func (m *customSmsCouponHistoryModel) Count(ctx context.Context) (int64, error) {
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

func (m *customSmsCouponHistoryModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
