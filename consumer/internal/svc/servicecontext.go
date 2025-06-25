package svc

import (
	"context"
	"fmt"
	"github.com/feihua/zero-admin/consumer/internal/config"
	consumer "github.com/feihua/zero-admin/consumer/internal/mq"
	"github.com/feihua/zero-admin/pkg/mq"
	"github.com/feihua/zero-admin/rpc/sms/client/couponrecordservice"
	"github.com/feihua/zero-admin/rpc/sms/client/couponservice"
	"github.com/feihua/zero-admin/rpc/sms/client/coupontypeservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membergrowthlogservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberinfoservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberpointslogservice"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	RabbitMQ *mq.RabbitMQ

	// 会员相关
	MemberInfoService      memberinfoservice.MemberInfoService
	MemberGrowthLogService membergrowthlogservice.MemberGrowthLogService
	MemberPointsLogService memberpointslogservice.MemberPointsLogService

	// 营销相关
	CouponRecordService couponrecordservice.CouponRecordService
	CouponService       couponservice.CouponService
	CouponTypeService   coupontypeservice.CouponTypeService
}

func NewServiceContext(c config.Config) *ServiceContext {
	umsClient := zrpc.MustNewClient(c.UmsRpc)
	smsClient := zrpc.MustNewClient(c.SmsRpc)

	mqUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Rabbitmq.UserName, c.Rabbitmq.Password, c.Rabbitmq.Host, c.Rabbitmq.Port)
	rabbitmq := mq.NewRabbitMQSimple(mqUrl)

	memberInfoService := memberinfoservice.NewMemberInfoService(umsClient)
	couponService := couponservice.NewCouponService(smsClient)
	couponRecordService := couponrecordservice.NewCouponRecordService(smsClient)
	s := &ServiceContext{
		Config:                 c,
		RabbitMQ:               rabbitmq,
		MemberInfoService:      memberInfoService,
		MemberGrowthLogService: membergrowthlogservice.NewMemberGrowthLogService(umsClient),
		MemberPointsLogService: memberpointslogservice.NewMemberPointsLogService(umsClient),
		CouponRecordService:    couponRecordService,
		CouponService:          couponService,
		CouponTypeService:      coupontypeservice.NewCouponTypeService(smsClient),
	}

	go func() {
		rabbitmq.ConsumeSimple("test", func(body []byte) {
			logc.Infof(context.Background(), "收到消息: %s", body)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("first_login_queue", func(body []byte) {
			consumer.FirstLogin(context.Background(), body, memberInfoService, couponService, couponRecordService)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("order.cancel.queue", func(body []byte) {
			logc.Infof(context.Background(), "收到order.cancel.queue消息: %s", body)
		})
	}()
	return s
}
