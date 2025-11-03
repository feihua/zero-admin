package svc

import (
	"time"

	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	Redis    *redis.Redis
	RedisKey string // redis的模块统一前缀
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisKey := c.RpcServerConf.Redis.Key
	rds := redis.MustNewRedis(c.RpcServerConf.Redis.RedisConf)

	db, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 settingLogConfig(),
	})
	if err != nil {
		logx.Errorf("mysql连接失败：%s", err.Error())
		panic(err)
	}

	logx.Debug("mysql已连接")
	query.SetDefault(db)

	return &ServiceContext{
		Config:   c,
		DB:       db,
		Redis:    rds,
		RedisKey: redisKey,
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
