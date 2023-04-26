package pmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmsAlbumPicModel = (*customPmsAlbumPicModel)(nil)

type (
	// PmsAlbumPicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsAlbumPicModel.
	PmsAlbumPicModel interface {
		pmsAlbumPicModel
	}

	customPmsAlbumPicModel struct {
		*defaultPmsAlbumPicModel
	}
)

// NewPmsAlbumPicModel returns a model for the database table.
func NewPmsAlbumPicModel(conn sqlx.SqlConn) PmsAlbumPicModel {
	return &customPmsAlbumPicModel{
		defaultPmsAlbumPicModel: newPmsAlbumPicModel(conn),
	}
}
