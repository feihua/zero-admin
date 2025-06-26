package mq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zeromicro/go-zero/core/logx"
)

// RabbitMQ 配置
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	MqUrl   string
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(MqUrl string) *RabbitMQ {
	return &RabbitMQ{MqUrl: MqUrl}
}

// Destroy 断开channel 和 connection
func (r *RabbitMQ) Destroy() {
	err := r.channel.Close()
	if err != nil {
		return
	}
	err = r.conn.Close()
	if err != nil {
		return
	}
}

// NewRabbitMQSimple 创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(MqUrl string) *RabbitMQ {
	// 创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(MqUrl)
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	if err != nil {
		logx.Errorf("rabbitmq获取connection失败：%s:%+v", "failed to connect rabbitmq!", err)
		panic(err)
	}
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		logx.Errorf("rabbitmq获取channel失败：%s:%+v", "failed to open a channel", err)
		panic(err)
	}

	logx.Info("rabbitmq连接成功")
	return rabbitmq
}

// PublishSimple 直接模式队列生产
func (r *RabbitMQ) PublishSimple(queueName string, message []byte) error {
	// 1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		queueName,
		// 是否持久化
		true,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞处理
		false,
		// 额外的属性
		nil,
	)
	if err != nil {
		r.Destroy()
		logx.Errorf("rabbitmq申请队列：%s失败, 错误消息: %+v", queueName, err)
		return fmt.Errorf("rabbitmq申请队列失败: %v", err)
	}
	// 调用channel 发送消息到队列中
	return r.channel.Publish(
		"",
		queueName,
		// 如果为true，根据自身exchange类型和routeKey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		// 如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
}

// ConsumeSimple simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple(queueName string, handler func([]byte)) {
	// 1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		queueName,
		// 是否持久化
		true,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞处理
		false,
		// 额外的属性
		nil,
	)
	if err != nil {
		r.Destroy()
		logx.Errorf("rabbitmq申请队列：%s失败, 错误消息: %+v", queueName, err)
		panic(err)
	}

	// 接收消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		// 用来区分多个消费者
		"", // consumer
		// 是否自动应答
		true, // auto-ack
		// 是否独有
		false, // exclusive
		// 设置为true，表示 不能将同一个Connection中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		// 列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		logx.Errorf("rabbitmq队列：%s接收消息失败, 错误消息: %+v", queueName, err)
		panic(err)
	}

	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			// 消息逻辑处理
			// log.Printf("Received a message: %s", d.Body)
			handler(d.Body)

		}
	}()

	logx.Info(" Waiting for messages ...")
	<-forever

}

// SendDelayMessage 发送延时取消消息
func (r *RabbitMQ) SendDelayMessage(exchange, queueName, key string, message []byte, delayMinutes int) error {
	// 声明延时交换机
	err := r.channel.ExchangeDeclare(
		exchange,            // 交换机名称
		"x-delayed-message", // 类型（需要延时插件）
		true,                // 持久化
		false,               // 自动删除
		false,               // 内部
		false,               // 不等待
		amqp.Table{
			"x-delayed-type": "direct",
		},
	)
	if err != nil {
		r.Destroy()
		logx.Errorf("声明延时交换机失败：%+v", err)
		return fmt.Errorf("声明延时交换机失败: %v", err)
	}

	// 声明订单取消队列
	_, err = r.channel.QueueDeclare(
		queueName, // 队列名称
		true,      // 持久化
		false,     // 自动删除
		false,     // 排他性
		false,     // 不等待
		nil,
	)
	if err != nil {
		r.Destroy()
		logx.Errorf("声明队列失败：%+v", err)
		return fmt.Errorf("声明队列失败: %v", err)
	}

	// 绑定队列到交换机
	err = r.channel.QueueBind(
		queueName, // 队列名
		key,       // 路由键
		exchange,  // 交换机
		false,
		nil,
	)
	if err != nil {
		r.Destroy()
		logx.Errorf("绑定队列失败：%+v", err)
		return fmt.Errorf("绑定队列失败: %v", err)
	}

	err = r.channel.Publish(
		exchange, // 交换机
		key,      // 路由键
		false,    // 强制
		false,    // 立即
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         message,
			DeliveryMode: amqp.Persistent,
			Headers: amqp.Table{
				"x-delay": int64(delayMinutes * 60 * 1000), // 延时时间（毫秒）
			},
		},
	)
	if err != nil {
		return fmt.Errorf("发送延时消息失败: %v", err)
	}

	return nil
}
