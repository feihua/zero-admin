package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-admin/front-api/internal/config"
	"zero-admin/rpc/oms/omsclient"
	"zero-admin/rpc/pay/payclient"
	"zero-admin/rpc/pms/pmsclient"
	"zero-admin/rpc/sms/smsclient"
	"zero-admin/rpc/sys/sysclient"
	"zero-admin/rpc/ums/umsclient"
)

type ServiceContext struct {
	Config config.Config

	Sys sysclient.Sys
	Ums umsclient.Ums
	Pms pmsclient.Pms
	Oms omsclient.Oms
	Sms smsclient.Sms
	Pay payclient.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Sys: sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
		Ums: umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc)),
		Pms: pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		Oms: omsclient.NewOms(zrpc.MustNewClient(c.OmsRpc)),
		Sms: smsclient.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		Pay: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
