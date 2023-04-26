package pmsmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PmsAlbumModel = (*customPmsAlbumModel)(nil)

type (
	// PmsAlbumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsAlbumModel.
	PmsAlbumModel interface {
		pmsAlbumModel
	}

	customPmsAlbumModel struct {
		*defaultPmsAlbumModel
	}
)

// NewPmsAlbumModel returns a model for the database table.
func NewPmsAlbumModel(conn sqlx.SqlConn) PmsAlbumModel {
	return &customPmsAlbumModel{
		defaultPmsAlbumModel: newPmsAlbumModel(conn),
	}
}
