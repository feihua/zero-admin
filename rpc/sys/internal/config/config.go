package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Postgresql struct {
		Datasource string
	}

	JWT struct {
		AccessSecret string
		AccessExpire int64
	}
}
