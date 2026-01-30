package couponservicelogic

import (
	"context"
	"errors"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// HandleExpirationCouponLogic 每天凌晨处理过期的优惠券
/*
Author: LiuFeiHua
Date: 2026/1/30 14:55
*/
type HandleExpirationCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandleExpirationCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleExpirationCouponLogic {
	return &HandleExpirationCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HandleExpirationCoupon 处理过期的优惠券
func (l *HandleExpirationCouponLogic) HandleExpirationCoupon(in *smsclient.HandleExpirationCouponReq) (*smsclient.HandleExpirationCouponResp, error) {

	err := l.svcCtx.DB.Exec(`
    UPDATE sms_coupon_record 
    SET status = 2 
    WHERE status = 0 
    AND coupon_id IN (
        SELECT id FROM sms_coupon WHERE end_time <= ?
    )`, time.Now()).Error

	if err != nil {
		logc.Errorf(l.ctx, "处理过期的优惠券,异常:%s", err.Error())
		return nil, errors.New("处理过期的优惠券异常")
	}

	return &smsclient.HandleExpirationCouponResp{}, nil
}
