package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		Datasource string
	}

	JWT struct {
		AccessSecret string
		AccessExpire int64
	}
}
