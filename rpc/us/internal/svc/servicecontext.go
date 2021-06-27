package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"go-zero-admin/rpc/model/usmodel"
	"go-zero-admin/rpc/us/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	UsUsersModel usmodel.UsUsersModel
	UsRolesModel usmodel.UsRolesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.Datasource)
	usUsersModel := usmodel.NewUsUsersModel(conn, c.CacheRedis)
	usRolesModel := usmodel.NewUsRolesModel(conn, c.CacheRedis)
	return &ServiceContext{
		Config:       c,
		UsUsersModel: usUsersModel,
		UsRolesModel: usRolesModel,
	}
}
