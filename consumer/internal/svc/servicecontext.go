package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/feihua/zero-admin/consumer/internal/config"
	"github.com/feihua/zero-admin/pkg/mq"
	"github.com/zeromicro/go-zero/core/logc"
	"log"
)

type ServiceContext struct {
	Config   config.Config
	RabbitMQ *mq.RabbitMQ
}

func NewServiceContext(c config.Config) *ServiceContext {
	mqUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Rabbitmq.UserName, c.Rabbitmq.Password, c.Rabbitmq.Host, c.Rabbitmq.Port)
	rabbitmq := mq.NewRabbitMQSimple(mqUrl)

	rabbitmq.ConsumeSimple("test", func(body []byte) {
		log.Printf("收到消息: %s", body)
		var msg map[string]interface{}
		err := json.Unmarshal(body, &msg)
		if err != nil {
			logc.Errorf(context.Background(), "序列化 JSON 失败: %v", err)
		}

		log.Printf("消息内容: %v", msg["nickname"])
	})
	return &ServiceContext{
		Config:   c,
		RabbitMQ: rabbitmq,
	}
}
