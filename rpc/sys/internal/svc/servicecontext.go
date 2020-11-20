package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"go-zero-admin/rpc/model"
	"go-zero-admin/rpc/sys/internal/config"
)

type ServiceContext struct {
	c config.Config

	UserModel     *model.SysUserModel
	RoleModel     *model.SysRoleModel
	MenuModel     *model.SysMenuModel
	DictModel     *model.SysDictModel
	DeptModel     *model.SysDeptModel
	LoginLogModel *model.SysLoginLogModel
	SysLogModel   *model.SysLogModel
	ConfigModel   *model.SysConfigModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		c:             c,
		UserModel:     model.NewSysUserModel(sqlConn),
		RoleModel:     model.NewSysRoleModel(sqlConn),
		MenuModel:     model.NewSysMenuModel(sqlConn),
		DictModel:     model.NewSysDictModel(sqlConn),
		DeptModel:     model.NewSysDeptModel(sqlConn),
		LoginLogModel: model.NewSysLoginLogModel(sqlConn),
		SysLogModel:   model.NewSysLogModel(sqlConn),
		ConfigModel:   model.NewSysConfigModel(sqlConn),
	}
}
