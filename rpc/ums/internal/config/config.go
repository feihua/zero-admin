package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		Datasource string
	}

	Mongo struct {
		Datasource string
		Db         string
	}
	JWT struct {
		AccessSecret string
		AccessExpire int64
	}
}
