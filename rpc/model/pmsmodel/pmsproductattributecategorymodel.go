package pmsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsProductAttributeCategoryModel = (*customPmsProductAttributeCategoryModel)(nil)

type (
	// PmsProductAttributeCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductAttributeCategoryModel.
	PmsProductAttributeCategoryModel interface {
		pmsProductAttributeCategoryModel
	}

	customPmsProductAttributeCategoryModel struct {
		*defaultPmsProductAttributeCategoryModel
	}
)

// NewPmsProductAttributeCategoryModel returns a model for the database table.
func NewPmsProductAttributeCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) PmsProductAttributeCategoryModel {
	return &customPmsProductAttributeCategoryModel{
		defaultPmsProductAttributeCategoryModel: newPmsProductAttributeCategoryModel(conn, c),
	}
}
