package pmsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsSkuStockModel = (*customPmsSkuStockModel)(nil)

type (
	// PmsSkuStockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsSkuStockModel.
	PmsSkuStockModel interface {
		pmsSkuStockModel
	}

	customPmsSkuStockModel struct {
		*defaultPmsSkuStockModel
	}
)

// NewPmsSkuStockModel returns a model for the database table.
func NewPmsSkuStockModel(conn sqlx.SqlConn, c cache.CacheConf) PmsSkuStockModel {
	return &customPmsSkuStockModel{
		defaultPmsSkuStockModel: newPmsSkuStockModel(conn, c),
	}
}
