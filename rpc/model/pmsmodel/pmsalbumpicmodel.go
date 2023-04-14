package pmsmodel

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsAlbumPicModel = (*customPmsAlbumPicModel)(nil)

type (
	// PmsAlbumPicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsAlbumPicModel.
	PmsAlbumPicModel interface {
		pmsAlbumPicModel
		FindListByAlbumId(ctx context.Context, albumId int64, Current int64, PageSize int64) (*[]PmsAlbumPic, error)
		CountByAlbumId(ctx context.Context, albumId int64) (int64, error)
	}

	customPmsAlbumPicModel struct {
		*defaultPmsAlbumPicModel
	}
)

// NewPmsAlbumPicModel returns a model for the database table.
func NewPmsAlbumPicModel(conn sqlx.SqlConn, c cache.CacheConf) PmsAlbumPicModel {
	return &customPmsAlbumPicModel{
		defaultPmsAlbumPicModel: newPmsAlbumPicModel(conn, c),
	}
}

func (m *customPmsAlbumPicModel) FindListByAlbumId(ctx context.Context, albumId int64, Current int64, PageSize int64) (*[]PmsAlbumPic, error) {

	query := fmt.Sprintf("select %s from %s where album_id=? limit ?,?", pmsAlbumRows, m.table)
	var resp []PmsAlbumPic
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, albumId, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPmsAlbumPicModel) CountByAlbumId(ctx context.Context, albumId int64) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where album_id=?", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query, albumId)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
