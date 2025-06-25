package svc

import (
	"fmt"
	"github.com/feihua/zero-admin/consumer/internal/config"
	"github.com/feihua/zero-admin/pkg/mq"
)

type ServiceContext struct {
	Config   config.Config
	RabbitMQ *mq.RabbitMQ
}

func NewServiceContext(c config.Config) *ServiceContext {
	mqUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Rabbitmq.UserName, c.Rabbitmq.Password, c.Rabbitmq.Host, c.Rabbitmq.Port)
	rabbitmq := mq.NewRabbitMQSimple(mqUrl)

	rabbitmq.ConsumeSimple("test")
	return &ServiceContext{
		Config:   c,
		RabbitMQ: rabbitmq,
	}
}
