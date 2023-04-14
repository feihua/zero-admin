package pmsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsProductCategoryAttributeRelationModel = (*customPmsProductCategoryAttributeRelationModel)(nil)

type (
	// PmsProductCategoryAttributeRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsProductCategoryAttributeRelationModel.
	PmsProductCategoryAttributeRelationModel interface {
		pmsProductCategoryAttributeRelationModel
	}

	customPmsProductCategoryAttributeRelationModel struct {
		*defaultPmsProductCategoryAttributeRelationModel
	}
)

// NewPmsProductCategoryAttributeRelationModel returns a model for the database table.
func NewPmsProductCategoryAttributeRelationModel(conn sqlx.SqlConn, c cache.CacheConf) PmsProductCategoryAttributeRelationModel {
	return &customPmsProductCategoryAttributeRelationModel{
		defaultPmsProductCategoryAttributeRelationModel: newPmsProductCategoryAttributeRelationModel(conn, c),
	}
}
