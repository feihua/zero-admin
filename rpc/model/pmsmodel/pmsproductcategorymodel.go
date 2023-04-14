package pmsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsProductCategoryModel = (*customPmsProductCategoryModel)(nil)

type (
	// PmsProductCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductCategoryModel.
	PmsProductCategoryModel interface {
		pmsProductCategoryModel
	}

	customPmsProductCategoryModel struct {
		*defaultPmsProductCategoryModel
	}
)

// NewPmsProductCategoryModel returns a model for the database table.
func NewPmsProductCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) PmsProductCategoryModel {
	return &customPmsProductCategoryModel{
		defaultPmsProductCategoryModel: newPmsProductCategoryModel(conn, c),
	}
}
