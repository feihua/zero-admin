package svc

import (
	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UmsGrowthChangeHistoryModel           umsmodel.UmsGrowthChangeHistoryModel
	UmsIntegrationChangeHistoryModel      umsmodel.UmsIntegrationChangeHistoryModel
	UmsIntegrationConsumeSettingModel     umsmodel.UmsIntegrationConsumeSettingModel
	UmsMemberModel                        umsmodel.UmsMemberModel
	UmsMemberAuthModel                    umsmodel.UmsMemberAuthModel
	UmsMemberLevelModel                   umsmodel.UmsMemberLevelModel
	UmsMemberLoginLogModel                umsmodel.UmsMemberLoginLogModel
	UmsMemberMemberTagRelationModel       umsmodel.UmsMemberMemberTagRelationModel
	UmsMemberPrepaidCardModel             umsmodel.UmsMemberPrepaidCardModel
	UmsMemberPrepaidCardLogModel          umsmodel.UmsMemberPrepaidCardLogModel
	UmsMemberPrepaidCardRelationModel     umsmodel.UmsMemberPrepaidCardRelationModel
	UmsMemberProductCategoryRelationModel umsmodel.UmsMemberProductCategoryRelationModel
	UmsMemberReceiveAddressModel          umsmodel.UmsMemberReceiveAddressModel
	UmsMemberRuleSettingModel             umsmodel.UmsMemberRuleSettingModel
	UmsMemberStatisticsInfoModel          umsmodel.UmsMemberStatisticsInfoModel
	UmsMemberTaskModel                    umsmodel.UmsMemberTaskModel
	UmsMemberTagModel                     umsmodel.UmsMemberTagModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config: c,

		UmsGrowthChangeHistoryModel:           umsmodel.NewUmsGrowthChangeHistoryModel(sqlConn, c.Cache),
		UmsIntegrationChangeHistoryModel:      umsmodel.NewUmsIntegrationChangeHistoryModel(sqlConn, c.Cache),
		UmsIntegrationConsumeSettingModel:     umsmodel.NewUmsIntegrationConsumeSettingModel(sqlConn, c.Cache),
		UmsMemberModel:                        umsmodel.NewUmsMemberModel(sqlConn, c.Cache),
		UmsMemberAuthModel:                    umsmodel.NewUmsMemberAuthModel(sqlConn, c.Cache),
		UmsMemberLevelModel:                   umsmodel.NewUmsMemberLevelModel(sqlConn, c.Cache),
		UmsMemberLoginLogModel:                umsmodel.NewUmsMemberLoginLogModel(sqlConn, c.Cache),
		UmsMemberMemberTagRelationModel:       umsmodel.NewUmsMemberMemberTagRelationModel(sqlConn, c.Cache),
		UmsMemberPrepaidCardModel:             umsmodel.NewUmsMemberPrepaidCardModel(sqlConn, c.Cache),
		UmsMemberPrepaidCardLogModel:          umsmodel.NewUmsMemberPrepaidCardLogModel(sqlConn, c.Cache),
		UmsMemberPrepaidCardRelationModel:     umsmodel.NewUmsMemberPrepaidCardRelationModel(sqlConn, c.Cache),
		UmsMemberProductCategoryRelationModel: umsmodel.NewUmsMemberProductCategoryRelationModel(sqlConn, c.Cache),
		UmsMemberReceiveAddressModel:          umsmodel.NewUmsMemberReceiveAddressModel(sqlConn, c.Cache),
		UmsMemberRuleSettingModel:             umsmodel.NewUmsMemberRuleSettingModel(sqlConn, c.Cache),
		UmsMemberStatisticsInfoModel:          umsmodel.NewUmsMemberStatisticsInfoModel(sqlConn, c.Cache),
		UmsMemberTaskModel:                    umsmodel.NewUmsMemberTaskModel(sqlConn, c.Cache),
		UmsMemberTagModel:                     umsmodel.NewUmsMemberTagModel(sqlConn, c.Cache),
	}
}
