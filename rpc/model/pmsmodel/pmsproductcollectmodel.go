package pmsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PmsProductCollectModel = (*customPmsProductCollectModel)(nil)

type (
	// PmsProductCollectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductCollectModel.
	PmsProductCollectModel interface {
		pmsProductCollectModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PmsProductCollect, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customPmsProductCollectModel struct {
		*defaultPmsProductCollectModel
	}
)

// NewPmsProductCollectModel returns a model for the database table.
func NewPmsProductCollectModel(conn sqlx.SqlConn) PmsProductCollectModel {
	return &customPmsProductCollectModel{
		defaultPmsProductCollectModel: newPmsProductCollectModel(conn),
	}
}

func (m *customPmsProductCollectModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PmsProductCollect, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsProductCollectRows, m.table)
	var resp []PmsProductCollect
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

func (m *customPmsProductCollectModel) Count(ctx context.Context) (int64, error) {
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

func (m *customPmsProductCollectModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
