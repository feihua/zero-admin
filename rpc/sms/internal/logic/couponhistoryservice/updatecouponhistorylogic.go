package couponhistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
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
	now := time.Now()
	_, err := q.WithContext(l.ctx).Updates(&model.SmsCouponHistory{
		ID:             in.Id,             //
		CouponID:       in.CouponId,       // 优惠券id
		MemberID:       in.MemberId,       // 会员id
		CouponCode:     in.CouponCode,     // 优惠码
		MemberNickname: in.MemberNickname, // 领取人昵称
		GetType:        in.GetType,        // 获取类型：0->后台赠送；1->主动获取
		UseStatus:      in.UseStatus,      // 使用状态：0->未使用；1->已使用；2->已过期
		UseTime:        &now,              // 使用时间
		OrderID:        &in.OrderId,       // 订单编号
		OrderSn:        &in.OrderSn,       // 订单号码
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券使用失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新优惠券使用失败")
	}

	return &smsclient.UpdateCouponHistoryResp{}, nil
}
