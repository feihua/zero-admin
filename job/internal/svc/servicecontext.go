package svc

import (
	"context"
	"fmt"
	"github.com/feihua/zero-admin/job/internal/config"
	"github.com/feihua/zero-admin/job/internal/jobs"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/client/orderservice"
	"github.com/feihua/zero-admin/rpc/oms/client/ordersettingservice"
	"github.com/feihua/zero-admin/rpc/pms/client/productskuservice"
	"github.com/feihua/zero-admin/rpc/sms/client/couponrecordservice"
	"github.com/feihua/zero-admin/rpc/sms/client/couponservice"
	"github.com/feihua/zero-admin/rpc/sms/client/coupontypeservice"
	"github.com/feihua/zero-admin/rpc/ums/client/membergrowthlogservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberinfoservice"
	"github.com/feihua/zero-admin/rpc/ums/client/memberpointslogservice"
	"github.com/go-co-op/gocron/v2"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type ServiceContext struct {
	Config config.Config

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
	// 订单相关
	OrderService        orderservice.OrderService
	OrderSettingService ordersettingservice.OrderSettingService
}

func NewServiceContext(c config.Config) *ServiceContext {
	umsClient := zrpc.MustNewClient(c.UmsRpc)
	smsClient := zrpc.MustNewClient(c.SmsRpc)
	omsClient := zrpc.MustNewClient(c.OmsRpc)
	pmsClient := zrpc.MustNewClient(c.PmsRpc)

	memberInfoService := memberinfoservice.NewMemberInfoService(umsClient)
	couponService := couponservice.NewCouponService(smsClient)
	couponRecordService := couponrecordservice.NewCouponRecordService(smsClient)
	skuService := productskuservice.NewProductSkuService(pmsClient)
	orderService := orderservice.NewOrderService(omsClient)
	settingService := ordersettingservice.NewOrderSettingService(omsClient)
	svc := &ServiceContext{
		Config:                 c,
		MemberInfoService:      memberInfoService,
		MemberGrowthLogService: membergrowthlogservice.NewMemberGrowthLogService(umsClient),
		MemberPointsLogService: memberpointslogservice.NewMemberPointsLogService(umsClient),
		CouponRecordService:    couponRecordService,
		CouponService:          couponService,
		CouponTypeService:      coupontypeservice.NewCouponTypeService(smsClient),
		ProductSkuService:      skuService,
		OrderService:           orderService,
		OrderSettingService:    settingService,
	}

	s, _ := gocron.NewScheduler()

	// 每隔10秒执行一次任务
	_, _ = s.NewJob(
		gocron.DurationJob(time.Second*10),
		gocron.NewTask(task, "test"),
	)

	// 每隔60秒执行一次任务
	_, _ = s.NewJob(
		gocron.DurationJob(time.Second*60),
		gocron.NewTask(jobs.CancelTimeOutOrder, context.Background(), skuService, orderService, couponRecordService, memberInfoService, settingService),
	)

	s.Start()
	return svc
}

func task(a string) {
	fmt.Println("执行任务:", a, time_util.TimeToStr(time.Now()))
}
