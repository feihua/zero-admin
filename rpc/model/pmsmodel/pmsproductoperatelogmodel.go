package pmsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PmsProductOperateLogModel = (*customPmsProductOperateLogModel)(nil)

type (
	// PmsProductOperateLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductOperateLogModel.
	PmsProductOperateLogModel interface {
		pmsProductOperateLogModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PmsProductOperateLog, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customPmsProductOperateLogModel struct {
		*defaultPmsProductOperateLogModel
	}
)

// NewPmsProductOperateLogModel returns a model for the database table.
func NewPmsProductOperateLogModel(conn sqlx.SqlConn) PmsProductOperateLogModel {
	return &customPmsProductOperateLogModel{
		defaultPmsProductOperateLogModel: newPmsProductOperateLogModel(conn),
	}
}

func (m *customPmsProductOperateLogModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PmsProductOperateLog, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductOperateLogRows, m.table)
	var resp []PmsProductOperateLog
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

func (m *customPmsProductOperateLogModel) Count(ctx context.Context) (int64, error) {
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

func (m *customPmsProductOperateLogModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
