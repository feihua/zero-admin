package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Rabbitmq struct {
		Host     string `json:",env=rabbitmq_host"`
		Port     int64  `json:",env=rabbitmq_port"`
		UserName string `json:",env=rabbitmq_user"`
		Password string `json:",env=rabbitmq_pass"`
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
