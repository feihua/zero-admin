package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Postgresql struct {
		Datasource string `json:",env=postgresql_dns"`
	}
	Rabbitmq struct {
		Host     string `json:",env=rabbitmq_host"`
		Port     int64  `json:",env=rabbitmq_port"`
		UserName string `json:",env=rabbitmq_user"`
		Password string `json:",env=rabbitmq_pass"`
	}
	Cart struct {
		Timeout int `json:",env=cart_timeout"`
	}
}
