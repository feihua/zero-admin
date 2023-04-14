package pmsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsProductAttributeValueModel = (*customPmsProductAttributeValueModel)(nil)

type (
	// PmsProductAttributeValueModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductAttributeValueModel.
	PmsProductAttributeValueModel interface {
		pmsProductAttributeValueModel
	}

	customPmsProductAttributeValueModel struct {
		*defaultPmsProductAttributeValueModel
	}
)

// NewPmsProductAttributeValueModel returns a model for the database table.
func NewPmsProductAttributeValueModel(conn sqlx.SqlConn, c cache.CacheConf) PmsProductAttributeValueModel {
	return &customPmsProductAttributeValueModel{
		defaultPmsProductAttributeValueModel: newPmsProductAttributeValueModel(conn, c),
	}
}
