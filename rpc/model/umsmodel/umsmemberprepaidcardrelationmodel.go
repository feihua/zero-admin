package umsmodel

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsMemberPrepaidCardRelationModel = (*customUmsMemberPrepaidCardRelationModel)(nil)

type (
	// UmsMemberPrepaidCardRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberPrepaidCardRelationModel.
	UmsMemberPrepaidCardRelationModel interface {
		umsMemberPrepaidCardRelationModel
		FindByIdAndMemberId(ctx context.Context, id int64, memberId int64) (*UmsMemberPrepaidCardRelation, error)
		FindListByMemberId(ctx context.Context, memberId int64, Current int64, PageSize int64) (*[]UmsMemberPrepaidCardRelation, error)
		CountByMemberId(ctx context.Context, memberId int64) (int64, error)
	}

	customUmsMemberPrepaidCardRelationModel struct {
		*defaultUmsMemberPrepaidCardRelationModel
	}
)

// NewUmsMemberPrepaidCardRelationModel returns a model for the database table.
func NewUmsMemberPrepaidCardRelationModel(conn sqlx.SqlConn, c cache.CacheConf) UmsMemberPrepaidCardRelationModel {
	return &customUmsMemberPrepaidCardRelationModel{
		defaultUmsMemberPrepaidCardRelationModel: newUmsMemberPrepaidCardRelationModel(conn, c),
	}
}

func (m *customUmsMemberPrepaidCardRelationModel) FindByIdAndMemberId(ctx context.Context, id int64, memberId int64) (*UmsMemberPrepaidCardRelation, error) {
	query := fmt.Sprintf("select %s from %s where id = ? and member_id = ? limit 1", umsMemberPrepaidCardRelationRows, m.table)
	var resp UmsMemberPrepaidCardRelation
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, id, memberId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *customUmsMemberPrepaidCardRelationModel) FindListByMemberId(ctx context.Context, memberId int64, Current int64, PageSize int64) (*[]UmsMemberPrepaidCardRelation, error) {

	query := fmt.Sprintf("select %s from %s where member_id = ? limit ?,?", umsMemberPrepaidCardRelationRows, m.table)
	var resp []UmsMemberPrepaidCardRelation
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, memberId, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUmsMemberPrepaidCardRelationModel) CountByMemberId(ctx context.Context, memberId int64) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where member_id = ?", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
