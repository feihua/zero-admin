package pmsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PmsFeightTemplateModel = (*customPmsFeightTemplateModel)(nil)

type (
	// PmsFeightTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPmsFeightTemplateModel.
	PmsFeightTemplateModel interface {
		pmsFeightTemplateModel
	}

	customPmsFeightTemplateModel struct {
		*defaultPmsFeightTemplateModel
	}
)

// NewPmsFeightTemplateModel returns a model for the database table.
func NewPmsFeightTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) PmsFeightTemplateModel {
	return &customPmsFeightTemplateModel{
		defaultPmsFeightTemplateModel: newPmsFeightTemplateModel(conn, c),
	}
}
