package svc

import (
	"zero-admin/rpc/model/omsmodel"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/oms/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	c       config.Config
	DbEngin *gorm.DB

	OmsCartItemModel            omsmodel.OmsCartItemModel
	OmsCompanyAddressModel      omsmodel.OmsCompanyAddressModel
	OmsOrderItemModel           omsmodel.OmsOrderItemModel
	OmsOrderModel               omsmodel.OmsOrderModel
	OmsOrderOperateHistoryModel omsmodel.OmsOrderOperateHistoryModel
	OmsOrderReturnApplyModel    omsmodel.OmsOrderReturnApplyModel
	OmsOrderReturnReasonModel   omsmodel.OmsOrderReturnReasonModel
	OmsOrderSettingModel        omsmodel.OmsOrderSettingModel
	PmsProductModel             pmsmodel.PmsProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "ums_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	// 如果出错就GameOver了
	if err != nil {
		panic(err)
	}
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c:       c,
		DbEngin: db,

		OmsCartItemModel:            omsmodel.NewOmsCartItemModel(sqlConn, c.Cache),
		OmsCompanyAddressModel:      omsmodel.NewOmsCompanyAddressModel(sqlConn, c.Cache),
		OmsOrderItemModel:           omsmodel.NewOmsOrderItemModel(sqlConn, c.Cache),
		OmsOrderModel:               omsmodel.NewOmsOrderModel(sqlConn, c.Cache),
		OmsOrderOperateHistoryModel: omsmodel.NewOmsOrderOperateHistoryModel(sqlConn, c.Cache),
		OmsOrderReturnApplyModel:    omsmodel.NewOmsOrderReturnApplyModel(sqlConn, c.Cache),
		OmsOrderReturnReasonModel:   omsmodel.NewOmsOrderReturnReasonModel(sqlConn, c.Cache),
		OmsOrderSettingModel:        omsmodel.NewOmsOrderSettingModel(sqlConn, c.Cache),
		PmsProductModel:             pmsmodel.NewPmsProductModel(sqlConn, c.Cache),
	}
}
