package pmsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PmsBrandModel = (*customPmsBrandModel)(nil)

type (
	// PmsBrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsBrandModel.
	PmsBrandModel interface {
		pmsBrandModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PmsBrand, error)
		DeleteByIds(ctx context.Context, ids []int64) error
		FindAllByIds(ctx context.Context, ids []int64) (*[]PmsBrand, error)
	}

	customPmsBrandModel struct {
		*defaultPmsBrandModel
	}
)

// NewPmsBrandModel returns a model for the database table.
func NewPmsBrandModel(conn sqlx.SqlConn) PmsBrandModel {
	return &customPmsBrandModel{
		defaultPmsBrandModel: newPmsBrandModel(conn),
	}
}

func (m *customPmsBrandModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]PmsBrand, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsBrandRows, m.table)
	var resp []PmsBrand
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

func (m *customPmsBrandModel) Count(ctx context.Context) (int64, error) {
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

func (m *customPmsBrandModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}

func (m *customPmsBrandModel) FindAllByIds(ctx context.Context, ids []int64) (*[]PmsBrand, error) {
	query := fmt.Sprintf("select %s from %s where `id` in (?)", pmsBrandRows, m.table)
	var resp []PmsBrand
	err := m.conn.QueryRows(&resp, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
