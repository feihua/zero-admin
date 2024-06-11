package svc

import (
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	//UmsGrowthChangeHistoryModel           umsmodel.UmsGrowthChangeHistoryModel
	//UmsIntegrationChangeHistoryModel      umsmodel.UmsIntegrationChangeHistoryModel
	//UmsIntegrationConsumeSettingModel     umsmodel.UmsIntegrationConsumeSettingModel
	//UmsMemberModel                        umsmodel.UmsMemberModel
	//UmsMemberLevelModel                   umsmodel.UmsMemberLevelModel
	//UmsMemberLoginLogModel                umsmodel.UmsMemberLoginLogModel
	//UmsMemberMemberTagRelationModel       umsmodel.UmsMemberMemberTagRelationModel
	//UmsMemberProductCategoryRelationModel umsmodel.UmsMemberProductCategoryRelationModel
	//UmsMemberReceiveAddressModel          umsmodel.UmsMemberReceiveAddressModel
	//UmsMemberRuleSettingModel             umsmodel.UmsMemberRuleSettingModel
	//UmsMemberStatisticsInfoModel          umsmodel.UmsMemberStatisticsInfoModel
	//UmsMemberTaskModel                    umsmodel.UmsMemberTaskModel
	//UmsMemberTagModel                     umsmodel.UmsMemberTagModel
	//UmsMemberReadHistoryModel             umsmodel.UmsMemberReadHistoryModel
	//UmsMemberProductCollectionModel       umsmodel.UmsMemberProductCollectionModel
	//UmsMemberBrandAttentionModel          umsmodel.UmsMemberBrandAttentionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 settingLogConfig(),
	})
	if err != nil {
		panic(err)
	}

	logx.Debug("mysql已连接")
	query.SetDefault(DB)
	//sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config: c,

		//UmsGrowthChangeHistoryModel:           umsmodel.NewUmsGrowthChangeHistoryModel(sqlConn),
		//UmsIntegrationChangeHistoryModel:      umsmodel.NewUmsIntegrationChangeHistoryModel(sqlConn),
		//UmsIntegrationConsumeSettingModel:     umsmodel.NewUmsIntegrationConsumeSettingModel(sqlConn),
		//UmsMemberModel:                        umsmodel.NewUmsMemberModel(sqlConn),
		//UmsMemberLevelModel:                   umsmodel.NewUmsMemberLevelModel(sqlConn),
		//UmsMemberLoginLogModel:                umsmodel.NewUmsMemberLoginLogModel(sqlConn),
		//UmsMemberMemberTagRelationModel:       umsmodel.NewUmsMemberMemberTagRelationModel(sqlConn),
		//UmsMemberProductCategoryRelationModel: umsmodel.NewUmsMemberProductCategoryRelationModel(sqlConn),
		//UmsMemberReceiveAddressModel:          umsmodel.NewUmsMemberReceiveAddressModel(sqlConn),
		//UmsMemberRuleSettingModel:             umsmodel.NewUmsMemberRuleSettingModel(sqlConn),
		//UmsMemberStatisticsInfoModel:          umsmodel.NewUmsMemberStatisticsInfoModel(sqlConn),
		//UmsMemberTaskModel:                    umsmodel.NewUmsMemberTaskModel(sqlConn),
		//UmsMemberTagModel:                     umsmodel.NewUmsMemberTagModel(sqlConn),
		//UmsMemberReadHistoryModel:             umsmodel.NewUmsMemberReadHistoryModel(sqlConn),
		//UmsMemberProductCollectionModel:       umsmodel.NewUmsMemberProductCollectionModel(sqlConn),
		//UmsMemberBrandAttentionModel:          umsmodel.NewUmsMemberBrandAttentionModel(sqlConn),
	}
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Infof(format, args...)
}

// init log config
func settingLogConfig() logger.Interface {
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return newLogger
}
