package svc

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-admin/api/internal/config"
	"go-zero-admin/api/internal/middleware"
	"go-zero-admin/rpc/oms/omsclient"
	"go-zero-admin/rpc/pms/pmsclient"
	"go-zero-admin/rpc/sms/smsclient"
	"go-zero-admin/rpc/sys/sysclient"
	"go-zero-admin/rpc/ums/umsclient"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	Sys      sysclient.Sys
	Ums      umsclient.Ums
	Pms      pmsclient.Pms
	Oms      omsclient.Oms
	Sms      smsclient.Sms
	Redis    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig())
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware(newRedis).Handle,
		Sys:      sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
		Ums:      umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		Pms:      pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		Oms:      omsclient.NewOms(zrpc.MustNewClient(c.OmsRpc)),
		Sms:      smsclient.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		Redis:    newRedis,
	}
}

func redisConfig() redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
	}
}
