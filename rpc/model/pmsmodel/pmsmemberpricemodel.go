package pmsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsMemberPriceModel = (*customPmsMemberPriceModel)(nil)

type (
	// PmsMemberPriceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsMemberPriceModel.
	PmsMemberPriceModel interface {
		pmsMemberPriceModel
	}

	customPmsMemberPriceModel struct {
		*defaultPmsMemberPriceModel
	}
)

// NewPmsMemberPriceModel returns a model for the database table.
func NewPmsMemberPriceModel(conn sqlx.SqlConn, c cache.CacheConf) PmsMemberPriceModel {
	return &customPmsMemberPriceModel{
		defaultPmsMemberPriceModel: newPmsMemberPriceModel(conn, c),
	}
}
