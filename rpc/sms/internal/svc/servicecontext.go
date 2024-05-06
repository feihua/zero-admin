package svc

import (
	"github.com/feihua/zero-admin/rpc/model/smsmodel"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type ServiceContext struct {
	Config                                config.Config
	DB                                    *gorm.DB
	SmsCouponHistoryModel                 smsmodel.SmsCouponHistoryModel
	SmsCouponModel                        smsmodel.SmsCouponModel
	SmsCouponProductCategoryRelationModel smsmodel.SmsCouponProductCategoryRelationModel
	SmsCouponProductRelationModel         smsmodel.SmsCouponProductRelationModel
	SmsFlashPromotionLogModel             smsmodel.SmsFlashPromotionLogModel
	SmsFlashPromotionModel                smsmodel.SmsFlashPromotionModel
	SmsFlashPromotionProductRelationModel smsmodel.SmsFlashPromotionProductRelationModel
	SmsFlashPromotionSessionModel         smsmodel.SmsFlashPromotionSessionModel
	SmsHomeAdvertiseModel                 smsmodel.SmsHomeAdvertiseModel
	SmsHomeBrandModel                     smsmodel.SmsHomeBrandModel
	SmsHomeNewProductModel                smsmodel.SmsHomeNewProductModel
	SmsHomeRecommendProductModel          smsmodel.SmsHomeRecommendProductModel
	SmsHomeRecommendSubjectModel          smsmodel.SmsHomeRecommendSubjectModel
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

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:                                c,
		DB:                                    DB,
		SmsCouponHistoryModel:                 smsmodel.NewSmsCouponHistoryModel(sqlConn),
		SmsCouponModel:                        smsmodel.NewSmsCouponModel(sqlConn),
		SmsCouponProductCategoryRelationModel: smsmodel.NewSmsCouponProductCategoryRelationModel(sqlConn),
		SmsCouponProductRelationModel:         smsmodel.NewSmsCouponProductRelationModel(sqlConn),
		SmsFlashPromotionLogModel:             smsmodel.NewSmsFlashPromotionLogModel(sqlConn),
		SmsFlashPromotionModel:                smsmodel.NewSmsFlashPromotionModel(sqlConn),
		SmsFlashPromotionProductRelationModel: smsmodel.NewSmsFlashPromotionProductRelationModel(sqlConn),
		SmsFlashPromotionSessionModel:         smsmodel.NewSmsFlashPromotionSessionModel(sqlConn),
		SmsHomeAdvertiseModel:                 smsmodel.NewSmsHomeAdvertiseModel(sqlConn),
		SmsHomeBrandModel:                     smsmodel.NewSmsHomeBrandModel(sqlConn),
		SmsHomeNewProductModel:                smsmodel.NewSmsHomeNewProductModel(sqlConn),
		SmsHomeRecommendProductModel:          smsmodel.NewSmsHomeRecommendProductModel(sqlConn),
		SmsHomeRecommendSubjectModel:          smsmodel.NewSmsHomeRecommendSubjectModel(sqlConn),
	}
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Debugf(format, args)
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
