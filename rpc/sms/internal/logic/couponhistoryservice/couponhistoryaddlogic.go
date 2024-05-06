package couponhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponHistoryAddLogic
/*
Author: LiuFeiHua
Date: 2024/5/6 11:34
*/
type CouponHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryAddLogic {
	return &CouponHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponHistoryAdd 添加优惠券使用记录
func (l *CouponHistoryAddLogic) CouponHistoryAdd(in *smsclient.CouponHistoryAddReq) (*smsclient.CouponHistoryAddResp, error) {
	err := query.SmsCouponHistory.WithContext(l.ctx).Create(&model.SmsCouponHistory{
		CouponID:       in.CouponId,
		MemberID:       in.MemberId,
		CouponCode:     in.CouponCode,
		MemberNickname: in.MemberNickname,
		GetType:        in.GetType,
		CreateTime:     time.Now(),
		UseStatus:      in.UseStatus,
		UseTime:        time.Now(),
		OrderID:        in.OrderId,
		OrderSn:        in.OrderSn,
	})
	if err != nil {
		return nil, err
	}

	return &smsclient.CouponHistoryAddResp{}, nil
}
