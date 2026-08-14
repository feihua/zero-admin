package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Postgresql struct {
		Datasource string `json:",env=postgresql_dns"`
	}

	Mongo struct {
		Datasource string `json:",env=mongo_dns"`
		Db         string `json:",env=mongo_db"`
	}
	Rabbitmq struct {
		Host     string `json:",env=rabbitmq_host"`
		Port     int64  `json:",env=rabbitmq_port"`
		UserName string `json:",env=rabbitmq_user"`
		Password string `json:",env=rabbitmq_pass"`
	}
	JWT struct {
		AccessSecret string `json:",env=jwt_access_secret"`
		AccessExpire int64  `json:",env=jwt_access_expire"`
	}
}
