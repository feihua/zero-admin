package jobs

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/client/couponservice"
	"github.com/zeromicro/go-zero/core/logc"
)

// HandleExpiredCouponLogic 处理过期的优惠券
func HandleExpiredCouponLogic(ctx context.Context, service couponservice.CouponService) {

	_, err := service.HandleExpirationCoupon(ctx, &couponservice.HandleExpirationCouponReq{})
	if err != nil {
		logc.Errorf(ctx, "处理过期的优惠券失败,错误信息：%+v", err)
		return
	}

}
