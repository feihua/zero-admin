package svc

import (
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type ServiceContext struct {
	c  config.Config
	DB *gorm.DB

	//OmsCartItemModel            omsmodel.OmsCartItemModel
	//OmsCompanyAddressModel      omsmodel.OmsCompanyAddressModel
	//OmsOrderItemModel           omsmodel.OmsOrderItemModel
	//OmsOrderModel               omsmodel.OmsOrderModel
	//OmsOrderOperateHistoryModel omsmodel.OmsOrderOperateHistoryModel
	//OmsOrderReturnApplyModel    omsmodel.OmsOrderReturnApplyModel
	//OmsOrderReturnReasonModel   omsmodel.OmsOrderReturnReasonModel
	//OmsOrderSettingModel        omsmodel.OmsOrderSettingModel
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
		c:  c,
		DB: DB,

		//OmsCartItemModel:            omsmodel.NewOmsCartItemModel(sqlConn),
		//OmsCompanyAddressModel:      omsmodel.NewOmsCompanyAddressModel(sqlConn),
		//OmsOrderItemModel:           omsmodel.NewOmsOrderItemModel(sqlConn),
		//OmsOrderModel:               omsmodel.NewOmsOrderModel(sqlConn),
		//OmsOrderOperateHistoryModel: omsmodel.NewOmsOrderOperateHistoryModel(sqlConn),
		//OmsOrderReturnApplyModel:    omsmodel.NewOmsOrderReturnApplyModel(sqlConn),
		//OmsOrderReturnReasonModel:   omsmodel.NewOmsOrderReturnReasonModel(sqlConn),
		//OmsOrderSettingModel:        omsmodel.NewOmsOrderSettingModel(sqlConn),
	}
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Infof(format, args)
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
