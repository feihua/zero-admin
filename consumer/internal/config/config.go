package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Rabbitmq struct {
		Host     string
		Port     int64
		UserName string
		Password string
	}
}
