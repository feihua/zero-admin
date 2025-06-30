package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Rabbitmq struct {
		Host     string
		Port     int64
		UserName string
		Password string
	}

	// 会员
	UmsRpc zrpc.RpcClientConf
	// 商品
	PmsRpc zrpc.RpcClientConf
	// 订单
	OmsRpc zrpc.RpcClientConf
	// 营销
	SmsRpc zrpc.RpcClientConf

	// 搜索
	SearchRpc zrpc.RpcClientConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
