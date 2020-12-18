package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-admin/front-api/internal/config"
	"go-zero-admin/rpc/oms/omsclient"
	"go-zero-admin/rpc/pms/pmsclient"
	"go-zero-admin/rpc/sms/smsclient"
	"go-zero-admin/rpc/sys/sysclient"
	"go-zero-admin/rpc/ums/umsclient"
)

type ServiceContext struct {
	Config config.Config

	Sys sysclient.Sys
	Ums umsclient.Ums
	Pms pmsclient.Pms
	Oms omsclient.Oms
	Sms smsclient.Sms
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Sys: sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
		Ums: umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		Pms: pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		Oms: omsclient.NewOms(zrpc.MustNewClient(c.OmsRpc)),
		Sms: smsclient.NewSms(zrpc.MustNewClient(c.SmsRpc)),
	}
}
