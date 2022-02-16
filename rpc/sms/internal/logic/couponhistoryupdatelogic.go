package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/smsmodel"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryUpdateLogic {
	return &CouponHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponHistoryUpdateLogic) CouponHistoryUpdate(in *sms.CouponHistoryUpdateReq) (*sms.CouponHistoryUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	UseTime, _ := time.Parse("2006-01-02 15:04:05", in.UseTime)
	err := l.svcCtx.SmsCouponHistoryModel.Update(smsmodel.SmsCouponHistory{
		Id:             in.Id,
		CouponId:       in.CouponId,
		MemberId:       in.MemberId,
		CouponCode:     in.CouponCode,
		MemberNickname: in.MemberNickname,
		GetType:        in.GetType,
		CreateTime:     CreateTime,
		UseStatus:      in.UseStatus,
		UseTime:        UseTime,
		OrderId:        in.OrderId,
		OrderSn:        in.OrderSn,
	})
	if err != nil {
		return nil, err
	}

	return &sms.CouponHistoryUpdateResp{}, nil
}
