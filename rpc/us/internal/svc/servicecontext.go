package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"go-zero-admin/rpc/model/usmodel"
	"go-zero-admin/rpc/us/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	UsUsersModel   usmodel.UsUsersModel
	UsRolesModel   usmodel.UsRolesModel
	UsTeacherModel usmodel.UsTeacherModel
	UsStudentModel usmodel.UsStudentModel
	RedisConn      *redis.Redis
	Conn           sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.Datasource)
	usUsersModel := usmodel.NewUsUsersModel(conn, c.CacheRedis)
	usRolesModel := usmodel.NewUsRolesModel(conn, c.CacheRedis)
	usTeacherModel := usmodel.NewUsTeacherModel(conn, c.CacheRedis)
	usStuModel := usmodel.NewUsStudentModel(conn, c.CacheRedis)
	rcon := redis.NewRedis(c.CacheRedis[0].Host, c.CacheRedis[0].Type)
	return &ServiceContext{
		Config:         c,
		UsUsersModel:   usUsersModel,
		UsRolesModel:   usRolesModel,
		UsTeacherModel: usTeacherModel,
		UsStudentModel: usStuModel,
		RedisConn:      rcon,
		Conn:           conn,
	}
}
