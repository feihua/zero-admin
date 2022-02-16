package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/internal/config"
)

type ServiceContext struct {
	Config config.Config

	UserModel     *sysmodel.SysUserModel
	UserRoleModel *sysmodel.SysUserRoleModel
	RoleModel     *sysmodel.SysRoleModel
	RoleMenuModel *sysmodel.SysRoleMenuModel
	MenuModel     *sysmodel.SysMenuModel
	DictModel     *sysmodel.SysDictModel
	DeptModel     *sysmodel.SysDeptModel
	LoginLogModel *sysmodel.SysLoginLogModel
	SysLogModel   *sysmodel.SysLogModel
	ConfigModel   *sysmodel.SysConfigModel
	JobModel      sysmodel.SysJobModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		Config:        c,
		UserModel:     sysmodel.NewSysUserModel(sqlConn),
		UserRoleModel: sysmodel.NewSysUserRoleModel(sqlConn),
		RoleModel:     sysmodel.NewSysRoleModel(sqlConn),
		RoleMenuModel: sysmodel.NewSysRoleMenuModel(sqlConn),
		MenuModel:     sysmodel.NewSysMenuModel(sqlConn),
		DictModel:     sysmodel.NewSysDictModel(sqlConn),
		DeptModel:     sysmodel.NewSysDeptModel(sqlConn),
		LoginLogModel: sysmodel.NewSysLoginLogModel(sqlConn),
		SysLogModel:   sysmodel.NewSysLogModel(sqlConn),
		ConfigModel:   sysmodel.NewSysConfigModel(sqlConn),
		JobModel:      sysmodel.NewSysJobModel(sqlConn),
	}
}
