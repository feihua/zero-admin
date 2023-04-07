package umsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsMemberProductCategoryRelationModel = (*customUmsMemberProductCategoryRelationModel)(nil)

type (
	// UmsMemberProductCategoryRelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberProductCategoryRelationModel.
	UmsMemberProductCategoryRelationModel interface {
		umsMemberProductCategoryRelationModel
	}

	customUmsMemberProductCategoryRelationModel struct {
		*defaultUmsMemberProductCategoryRelationModel
	}
)

// NewUmsMemberProductCategoryRelationModel returns a model for the database table.
func NewUmsMemberProductCategoryRelationModel(conn sqlx.SqlConn, c cache.CacheConf) UmsMemberProductCategoryRelationModel {
	return &customUmsMemberProductCategoryRelationModel{
		defaultUmsMemberProductCategoryRelationModel: newUmsMemberProductCategoryRelationModel(conn, c),
	}
}
