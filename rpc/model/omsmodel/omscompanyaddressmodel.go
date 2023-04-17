package omsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OmsCompanyAddressModel = (*customOmsCompanyAddressModel)(nil)

type (
	// OmsCompanyAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOmsCompanyAddressModel.
	OmsCompanyAddressModel interface {
		omsCompanyAddressModel
	}

	customOmsCompanyAddressModel struct {
		*defaultOmsCompanyAddressModel
	}
)

// NewOmsCompanyAddressModel returns a model for the database table.
func NewOmsCompanyAddressModel(conn sqlx.SqlConn, c cache.CacheConf) OmsCompanyAddressModel {
	return &customOmsCompanyAddressModel{
		defaultOmsCompanyAddressModel: newOmsCompanyAddressModel(conn, c),
	}
}
