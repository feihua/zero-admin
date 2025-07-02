package svc

import (
	"context"
	"fmt"
	"github.com/feihua/zero-admin/consumer/internal/config"
	"github.com/feihua/zero-admin/consumer/internal/mq/coupon"
	"github.com/feihua/zero-admin/consumer/internal/mq/order"
	"github.com/feihua/zero-admin/consumer/internal/mq/product"
	"github.com/feihua/zero-admin/pkg/mq"
	"github.com/feihua/zero-admin/rpc/oms/client/orderservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productskuservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productspuservice"
	"github.com/feihua/zero-admin/rpc/search/search_client"
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

	// 商品相关
	ProductSkuService productskuservice.ProductSkuService
	ProductSpuService productspuservice.ProductSpuService
	// 订单相关
	OrderService orderservice.OrderService

	// 搜索相关
	Search search_client.Search
}

func NewServiceContext(c config.Config) *ServiceContext {
	umsClient := zrpc.MustNewClient(c.UmsRpc)
	smsClient := zrpc.MustNewClient(c.SmsRpc)
	omsClient := zrpc.MustNewClient(c.OmsRpc)
	pmsClient := zrpc.MustNewClient(c.PmsRpc)
	searchClient := zrpc.MustNewClient(c.SearchRpc)

	mqUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Rabbitmq.UserName, c.Rabbitmq.Password, c.Rabbitmq.Host, c.Rabbitmq.Port)
	rabbitmq := mq.NewRabbitMQSimple(mqUrl)

	memberInfoService := memberinfoservice.NewMemberInfoService(umsClient)
	couponService := couponservice.NewCouponService(smsClient)
	couponRecordService := couponrecordservice.NewCouponRecordService(smsClient)
	skuService := productskuservice.NewProductSkuService(pmsClient)
	spuService := productspuservice.NewProductSpuService(pmsClient)
	orderService := orderservice.NewOrderService(omsClient)
	search := search_client.NewSearch(searchClient)
	s := &ServiceContext{
		Config:                 c,
		RabbitMQ:               rabbitmq,
		MemberInfoService:      memberInfoService,
		MemberGrowthLogService: membergrowthlogservice.NewMemberGrowthLogService(umsClient),
		MemberPointsLogService: memberpointslogservice.NewMemberPointsLogService(umsClient),
		CouponRecordService:    couponRecordService,
		CouponService:          couponService,
		CouponTypeService:      coupontypeservice.NewCouponTypeService(smsClient),
		ProductSkuService:      skuService,
		ProductSpuService:      spuService,
		OrderService:           orderService,
		Search:                 search,
	}

	go func() {
		rabbitmq.ConsumeSimple("test", func(body []byte) {
			logc.Infof(context.Background(), "收到消息: %s", body)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("first.login.queue", func(body []byte) {
			coupon.FirstLogin(context.Background(), body, memberInfoService, couponService, couponRecordService)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("order.cancel.queue", func(body []byte) {
			order.OrderCancel(context.Background(), body, skuService, orderService, couponRecordService, memberInfoService)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("syn.product.to.es.queue", func(body []byte) {
			product.SynProductToEs(context.Background(), body, search, spuService)
		})
	}()
	go func() {
		rabbitmq.ConsumeSimple("delete.product.from.es.queue", func(body []byte) {
			product.DeleteProductFromEs(context.Background(), body, search, spuService)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("order.return.queue", func(body []byte) {
			order.OrderReturn(context.Background(), body)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("order.cancel.by.user.queue", func(body []byte) {
			order.OrderCancelByUser(context.Background(), body)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("order.close.queue", func(body []byte) {
			order.OrderClose(context.Background(), body)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("order.delivery.queue", func(body []byte) {
			order.OrderDelivery(context.Background(), body)
		})
	}()

	go func() {
		rabbitmq.ConsumeSimple("order.confirm.queue", func(body []byte) {
			order.OrderConfirm(context.Background(), body)
		})
	}()
	return s
}
