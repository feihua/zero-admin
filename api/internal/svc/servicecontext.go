package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-admin/api/internal/config"
	"go-zero-admin/rpc/sys/sysclient"
)

type ServiceContext struct {
	Config config.Config

	Sys sysclient.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Sys: sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
