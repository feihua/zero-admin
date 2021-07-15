package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	//系统
	SysRpc zrpc.RpcClientConf
	//会员
	UmsRpc zrpc.RpcClientConf
	//商品
	PmsRpc zrpc.RpcClientConf
	//订单
	OmsRpc zrpc.RpcClientConf
	//营销
	SmsRpc zrpc.RpcClientConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Redis struct {
		Address string
	}
}
