package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"go-zero-admin/rpc/model/paymodel"
	"go-zero-admin/rpc/pay/internal/config"
)

type ServiceContext struct {
	c             config.Config
	WxRecordModel paymodel.PayWxRecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c:             c,
		WxRecordModel: paymodel.NewPayWxRecordModel(sqlConn),
	}
}
