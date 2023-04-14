package pmsmodel

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsAlbumModel = (*customPmsAlbumModel)(nil)

type (
	// PmsAlbumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsAlbumModel.
	PmsAlbumModel interface {
		pmsAlbumModel
		FindList(ctx context.Context, Current int64, PageSize int64) (*[]PmsAlbum, error)
		Count(ctx context.Context) (int64, error)
	}

	customPmsAlbumModel struct {
		*defaultPmsAlbumModel
	}
)

// NewPmsAlbumModel returns a model for the database table.
func NewPmsAlbumModel(conn sqlx.SqlConn, c cache.CacheConf) PmsAlbumModel {
	return &customPmsAlbumModel{
		defaultPmsAlbumModel: newPmsAlbumModel(conn, c),
	}
}

func (m *customPmsAlbumModel) FindList(ctx context.Context, Current int64, PageSize int64) (*[]PmsAlbum, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", pmsAlbumRows, m.table)
	var resp []PmsAlbum
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPmsAlbumModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)
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
