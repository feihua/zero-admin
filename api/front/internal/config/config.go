package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// 系统
	SysRpc zrpc.RpcClientConf
	// 会员
	UmsRpc zrpc.RpcClientConf
	// 商品
	PmsRpc zrpc.RpcClientConf
	// 订单
	OmsRpc zrpc.RpcClientConf
	// 营销
	SmsRpc zrpc.RpcClientConf
	// 支付
	PayRpc zrpc.RpcClientConf
	// 内容相关
	CmsRpc zrpc.RpcClientConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	// 支付宝支付配置
	Alipay struct {
		AppId        string
		PrivateKey   string
		ServerDomain string
		NotifyURL    string
		IsProduction bool
	}

	Rabbitmq struct {
		Host     string
		Port     int64
		UserName string
		Password string
	}
}
