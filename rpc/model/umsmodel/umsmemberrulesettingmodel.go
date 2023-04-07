package umsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsMemberRuleSettingModel = (*customUmsMemberRuleSettingModel)(nil)

type (
	// UmsMemberRuleSettingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberRuleSettingModel.
	UmsMemberRuleSettingModel interface {
		umsMemberRuleSettingModel
	}

	customUmsMemberRuleSettingModel struct {
		*defaultUmsMemberRuleSettingModel
	}
)

// NewUmsMemberRuleSettingModel returns a model for the database table.
func NewUmsMemberRuleSettingModel(conn sqlx.SqlConn, c cache.CacheConf) UmsMemberRuleSettingModel {
	return &customUmsMemberRuleSettingModel{
		defaultUmsMemberRuleSettingModel: newUmsMemberRuleSettingModel(conn, c),
	}
}
