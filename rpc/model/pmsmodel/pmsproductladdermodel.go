package pmsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsProductLadderModel = (*customPmsProductLadderModel)(nil)

type (
	// PmsProductLadderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductLadderModel.
	PmsProductLadderModel interface {
		pmsProductLadderModel
	}

	customPmsProductLadderModel struct {
		*defaultPmsProductLadderModel
	}
)

// NewPmsProductLadderModel returns a model for the database table.
func NewPmsProductLadderModel(conn sqlx.SqlConn, c cache.CacheConf) PmsProductLadderModel {
	return &customPmsProductLadderModel{
		defaultPmsProductLadderModel: newPmsProductLadderModel(conn, c),
	}
}
