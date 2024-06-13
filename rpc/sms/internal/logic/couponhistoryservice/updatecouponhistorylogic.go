package couponhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCouponHistoryLogic 更新优惠券使用、领取历史表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:30
*/
type UpdateCouponHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCouponHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponHistoryLogic {
	return &UpdateCouponHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCouponHistory 更新优惠券使用、领取历史表
func (l *UpdateCouponHistoryLogic) UpdateCouponHistory(in *smsclient.UpdateCouponHistoryReq) (*smsclient.UpdateCouponHistoryResp, error) {
	q := query.SmsCouponHistory
	_, err := q.WithContext(l.ctx).Updates(&model.SmsCouponHistory{
		ID:             in.Id,
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

	return &smsclient.UpdateCouponHistoryResp{}, nil
}
