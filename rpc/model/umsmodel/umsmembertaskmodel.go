package umsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsMemberTaskModel = (*customUmsMemberTaskModel)(nil)

type (
	// UmsMemberTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberTaskModel.
	UmsMemberTaskModel interface {
		umsMemberTaskModel
	}

	customUmsMemberTaskModel struct {
		*defaultUmsMemberTaskModel
	}
)

// NewUmsMemberTaskModel returns a model for the database table.
func NewUmsMemberTaskModel(conn sqlx.SqlConn, c cache.CacheConf) UmsMemberTaskModel {
	return &customUmsMemberTaskModel{
		defaultUmsMemberTaskModel: newUmsMemberTaskModel(conn, c),
	}
}
