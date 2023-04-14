package pmsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsProductAttributeModel = (*customPmsProductAttributeModel)(nil)

type (
	// PmsProductAttributeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductAttributeModel.
	PmsProductAttributeModel interface {
		pmsProductAttributeModel
	}

	customPmsProductAttributeModel struct {
		*defaultPmsProductAttributeModel
	}
)

// NewPmsProductAttributeModel returns a model for the database table.
func NewPmsProductAttributeModel(conn sqlx.SqlConn, c cache.CacheConf) PmsProductAttributeModel {
	return &customPmsProductAttributeModel{
		defaultPmsProductAttributeModel: newPmsProductAttributeModel(conn, c),
	}
}
