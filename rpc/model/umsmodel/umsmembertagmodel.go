package umsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UmsMemberTagModel = (*customUmsMemberTagModel)(nil)

type (
	// UmsMemberTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberTagModel.
	UmsMemberTagModel interface {
		umsMemberTagModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]UmsMemberTag, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customUmsMemberTagModel struct {
		*defaultUmsMemberTagModel
	}
)

// NewUmsMemberTagModel returns a model for the database table.
func NewUmsMemberTagModel(conn sqlx.SqlConn) UmsMemberTagModel {
	return &customUmsMemberTagModel{
		defaultUmsMemberTagModel: newUmsMemberTagModel(conn),
	}
}

func (m *customUmsMemberTagModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]UmsMemberTag, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberTagRows, m.table)
	var resp []UmsMemberTag
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

func (m *customUmsMemberTagModel) Count(ctx context.Context) (int64, error) {
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

func (m *customUmsMemberTagModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
