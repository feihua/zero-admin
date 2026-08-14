package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Postgresql struct {
		Datasource string `json:",env=postgresql_dns"`
	}

	JWT struct {
		AccessSecret string `json:",env=jwt_access_secret"`
		AccessExpire int64  `json:",env=jwt_access_expire"`
	}
}
