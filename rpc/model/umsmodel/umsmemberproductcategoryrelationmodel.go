package umsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UmsMemberProductCategoryRelationModel = (*customUmsMemberProductCategoryRelationModel)(nil)

type (
	// UmsMemberProductCategoryRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberProductCategoryRelationModel.
	UmsMemberProductCategoryRelationModel interface {
		umsMemberProductCategoryRelationModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]UmsMemberProductCategoryRelation, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customUmsMemberProductCategoryRelationModel struct {
		*defaultUmsMemberProductCategoryRelationModel
	}
)

// NewUmsMemberProductCategoryRelationModel returns a model for the database table.
func NewUmsMemberProductCategoryRelationModel(conn sqlx.SqlConn) UmsMemberProductCategoryRelationModel {
	return &customUmsMemberProductCategoryRelationModel{
		defaultUmsMemberProductCategoryRelationModel: newUmsMemberProductCategoryRelationModel(conn),
	}
}

func (m *customUmsMemberProductCategoryRelationModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]UmsMemberProductCategoryRelation, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberProductCategoryRelationRows, m.table)
	var resp []UmsMemberProductCategoryRelation
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

func (m *customUmsMemberProductCategoryRelationModel) Count(ctx context.Context) (int64, error) {
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

func (m *customUmsMemberProductCategoryRelationModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
