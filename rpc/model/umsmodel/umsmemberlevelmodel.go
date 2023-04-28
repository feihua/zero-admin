package umsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UmsMemberLevelModel = (*customUmsMemberLevelModel)(nil)

type (
	// UmsMemberLevelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberLevelModel.
	UmsMemberLevelModel interface {
		umsMemberLevelModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]UmsMemberLevel, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customUmsMemberLevelModel struct {
		*defaultUmsMemberLevelModel
	}
)

// NewUmsMemberLevelModel returns a model for the database table.
func NewUmsMemberLevelModel(conn sqlx.SqlConn) UmsMemberLevelModel {
	return &customUmsMemberLevelModel{
		defaultUmsMemberLevelModel: newUmsMemberLevelModel(conn),
	}
}

func (m *customUmsMemberLevelModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]UmsMemberLevel, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberLevelRows, m.table)
	var resp []UmsMemberLevel
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

func (m *customUmsMemberLevelModel) Count(ctx context.Context) (int64, error) {
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

func (m *customUmsMemberLevelModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
