package svc

import (
	"context"
	"fmt"
	"github.com/feihua/zero-admin/consumer/internal/config"
	consumer "github.com/feihua/zero-admin/consumer/internal/mq"
	"github.com/feihua/zero-admin/pkg/mq"
	"github.com/zeromicro/go-zero/core/logc"
)

type ServiceContext struct {
	Config   config.Config
	RabbitMQ *mq.RabbitMQ
}

func NewServiceContext(c config.Config) *ServiceContext {
	mqUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Rabbitmq.UserName, c.Rabbitmq.Password, c.Rabbitmq.Host, c.Rabbitmq.Port)
	rabbitmq := mq.NewRabbitMQSimple(mqUrl)

	go func() {
		rabbitmq.ConsumeSimple("test", func(body []byte) {
			logc.Infof(context.Background(), "收到消息: %s", body)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("first_login_queue", func(body []byte) {
			consumer.FirstLogin(context.Background(), body)
		})
	}()

	return &ServiceContext{
		Config:   c,
		RabbitMQ: rabbitmq,
	}
}
