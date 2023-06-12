package pmsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PmsProductVertifyRecordModel = (*customPmsProductVertifyRecordModel)(nil)

type (
	// PmsProductVertifyRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductVertifyRecordModel.
	PmsProductVertifyRecordModel interface {
		pmsProductVertifyRecordModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PmsProductVertifyRecord, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customPmsProductVertifyRecordModel struct {
		*defaultPmsProductVertifyRecordModel
	}
)

// NewPmsProductVertifyRecordModel returns a model for the database table.
func NewPmsProductVertifyRecordModel(conn sqlx.SqlConn) PmsProductVertifyRecordModel {
	return &customPmsProductVertifyRecordModel{
		defaultPmsProductVertifyRecordModel: newPmsProductVertifyRecordModel(conn),
	}
}

func (m *customPmsProductVertifyRecordModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PmsProductVertifyRecord, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductVertifyRecordRows, m.table)
	var resp []PmsProductVertifyRecord
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

func (m *customPmsProductVertifyRecordModel) Count(ctx context.Context) (int64, error) {
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

func (m *customPmsProductVertifyRecordModel) DeleteByIds(ctx context.Context, ids []int64) error {
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
