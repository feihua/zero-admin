package umsmodel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UmsIntegrationConsumeSettingModel = (*customUmsIntegrationConsumeSettingModel)(nil)

type (
	// UmsIntegrationConsumeSettingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsIntegrationConsumeSettingModel.
	UmsIntegrationConsumeSettingModel interface {
		umsIntegrationConsumeSettingModel
	}

	customUmsIntegrationConsumeSettingModel struct {
		*defaultUmsIntegrationConsumeSettingModel
	}
)

// NewUmsIntegrationConsumeSettingModel returns a model for the database table.
func NewUmsIntegrationConsumeSettingModel(conn sqlx.SqlConn, c cache.CacheConf) UmsIntegrationConsumeSettingModel {
	return &customUmsIntegrationConsumeSettingModel{
		defaultUmsIntegrationConsumeSettingModel: newUmsIntegrationConsumeSettingModel(conn, c),
	}
}
