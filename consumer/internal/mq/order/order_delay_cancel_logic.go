package order

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/rpc/oms/client/orderservice"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/feihua/zero-admin/rpc/pms/client/productskuservice"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/client/couponrecordservice"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/feihua/zero-admin/rpc/ums/client/memberinfoservice"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
)

// OrderDelayCancel 处理30分钟未支付的订单
func OrderDelayCancel(ctx context.Context, body []byte, productSkuService productskuservice.ProductSkuService, orderService orderservice.OrderService, couponRecordService couponrecordservice.CouponRecordService, memberService memberinfoservice.MemberInfoService) {
	logc.Infof(ctx, "收到超过30分钟未付款的订单mq消息: %s", body)
	var orderInfo map[string]int64
	err := sonic.Unmarshal(body, &orderInfo)
	if err != nil {
		logc.Errorf(ctx, "序列化 JSON 失败: %v", err)
		return
	}

	orderId := orderInfo["orderId"]
	memberId := orderInfo["memberId"]

	// todo 暂时没有分布式事务
	// 1.查询订单是否存在
	// 2.修改订单状态
	resp, err := orderService.CancelOrder(ctx, &omsclient.CancelOrderReq{
		MemberId: memberId,
		OrderId:  orderId,
	})
	if err != nil {
		logc.Errorf(ctx, "查询订单是否存在失败,请求参数：%+v,错误信息：%+v", orderInfo, err)
		return
	}

	couponIds := resp.CouponIds
	integration := resp.Integration
	stockLockData := resp.Data

	var data []*pmsclient.UpdateSkuStockData
	for _, item := range stockLockData {
		data = append(data, &pmsclient.UpdateSkuStockData{
			Id:              item.ProductSkuId,
			ProductQuantity: item.ProductQuantity,
		})
	}
	// 3.释放库存
	_, err = productSkuService.ReleaseSkuStockLock(ctx, &pmsclient.UpdateSkuStockReq{
		Data: data,
	})
	if err != nil {
		logc.Errorf(ctx, "释放库存失败,请求参数：%+v,错误信息：%+v", orderInfo, err)
		return
	}

	// 4.如果使用优惠券,更新优惠券使用状态
	// 4.1修改sms_coupon_history表的use_status字段
	// 4.2记得修改sms_coupon的use_count字段,下单的时候要加1,取消订单的时候,要减1
	if len(couponIds) > 0 {
		_, err = couponRecordService.UpdateCouponRecord(ctx, &smsclient.UpdateCouponRecordReq{
			MemberId:  memberId,
			Status:    0,
			CouponIds: couponIds,
		})
		if err != nil {
			logc.Errorf(ctx, "更新优惠券使用状态失败,请求参数：%+v,错误信息：%+v", orderInfo, err)
			return
		}
	}

	// 5.返还使用积分
	member, _ := memberService.QueryMemberInfoDetail(ctx, &umsclient.QueryMemberInfoDetailReq{MemberId: memberId})
	i := member.Points + integration
	_, err = memberService.UpdateMemberPoints(ctx, &umsclient.UpdateMemberPointsReq{MemberId: memberId, Points: i})
	if err != nil {
		logc.Errorf(ctx, "返还使用积分失败,请求参数：%+v,错误信息：%+v", orderInfo, err)
		return
	}

	logc.Errorf(ctx, "取消用户：%d 未支付的订单：%d 成功", memberId, orderId)

}
