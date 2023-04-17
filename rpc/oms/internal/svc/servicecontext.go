package svc

import (
	"zero-admin/rpc/model/omsmodel"
	"zero-admin/rpc/oms/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c config.Config

	OmsCartItemModel            omsmodel.OmsCartItemModel
	OmsCompanyAddressModel      omsmodel.OmsCompanyAddressModel
	OmsOrderItemModel           omsmodel.OmsOrderItemModel
	OmsOrderModel               omsmodel.OmsOrderModel
	OmsOrderOperateHistoryModel omsmodel.OmsOrderOperateHistoryModel
	OmsOrderReturnApplyModel    omsmodel.OmsOrderReturnApplyModel
	OmsOrderReturnReasonModel   omsmodel.OmsOrderReturnReasonModel
	OmsOrderSettingModel        omsmodel.OmsOrderSettingModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c: c,

		OmsCartItemModel:            omsmodel.NewOmsCartItemModel(sqlConn, c.Cache),
		OmsCompanyAddressModel:      omsmodel.NewOmsCompanyAddressModel(sqlConn, c.Cache),
		OmsOrderItemModel:           omsmodel.NewOmsOrderItemModel(sqlConn, c.Cache),
		OmsOrderModel:               omsmodel.NewOmsOrderModel(sqlConn, c.Cache),
		OmsOrderOperateHistoryModel: omsmodel.NewOmsOrderOperateHistoryModel(sqlConn, c.Cache),
		OmsOrderReturnApplyModel:    omsmodel.NewOmsOrderReturnApplyModel(sqlConn, c.Cache),
		OmsOrderReturnReasonModel:   omsmodel.NewOmsOrderReturnReasonModel(sqlConn, c.Cache),
		OmsOrderSettingModel:        omsmodel.NewOmsOrderSettingModel(sqlConn, c.Cache),
	}
}
