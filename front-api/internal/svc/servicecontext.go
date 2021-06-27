package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-admin/front-api/internal/config"
	"go-zero-admin/rpc/oms/omsclient"
	"go-zero-admin/rpc/pay/payclient"
	"go-zero-admin/rpc/pms/pmsclient"
	"go-zero-admin/rpc/sms/smsclient"
	"go-zero-admin/rpc/sys/sysclient"
	"go-zero-admin/rpc/ums/umsclient"
	"go-zero-admin/rpc/us/usclient"
)

type ServiceContext struct {
	Config config.Config

	Sys sysclient.Sys
	Ums umsclient.Ums
	Pms pmsclient.Pms
	Oms omsclient.Oms
	Sms smsclient.Sms
	Pay payclient.Pay
	Us  usclient.Us
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
		Us:  usclient.NewUs(zrpc.MustNewClient(c.UsRpc)),
	}
}
