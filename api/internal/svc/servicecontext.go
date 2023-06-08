package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-admin/api/internal/config"
	"zero-admin/api/internal/middleware"
	"zero-admin/rpc/cms/cms"
	"zero-admin/rpc/oms/oms"
	"zero-admin/rpc/pms/pms"
	"zero-admin/rpc/sms/sms"
	"zero-admin/rpc/sys/sys"
	"zero-admin/rpc/ums/ums"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	AddLog   rest.Middleware
	Sys      sys.Sys
	Ums      ums.Ums
	Pms      pms.Pms
	Oms      oms.Oms
	Sms      sms.Sms
	Cms      cms.Cms
	Redis    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	newSys := sys.NewSys(zrpc.MustNewClient(c.SysRpc))
	return &ServiceContext{
		Config:   c,
		Sys:      newSys,
		CheckUrl: middleware.NewCheckUrlMiddleware(newRedis).Handle,
		AddLog:   middleware.NewAddLogMiddleware(newSys).Handle,
		Ums:      ums.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		Pms:      pms.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		Oms:      oms.NewOms(zrpc.MustNewClient(c.OmsRpc)),
		Sms:      sms.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		Cms:      cms.NewCms(zrpc.MustNewClient(c.CmsRpc)),
		Redis:    newRedis,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
