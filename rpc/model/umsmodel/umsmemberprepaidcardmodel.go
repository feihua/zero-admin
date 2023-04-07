package umsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsMemberPrepaidCardModel = (*customUmsMemberPrepaidCardModel)(nil)

type (
	// UmsMemberPrepaidCardModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberPrepaidCardModel.
	UmsMemberPrepaidCardModel interface {
		umsMemberPrepaidCardModel
	}

	customUmsMemberPrepaidCardModel struct {
		*defaultUmsMemberPrepaidCardModel
	}
)

// NewUmsMemberPrepaidCardModel returns a model for the database table.
func NewUmsMemberPrepaidCardModel(conn sqlx.SqlConn, c cache.CacheConf) UmsMemberPrepaidCardModel {
	return &customUmsMemberPrepaidCardModel{
		defaultUmsMemberPrepaidCardModel: newUmsMemberPrepaidCardModel(conn, c),
	}
}
