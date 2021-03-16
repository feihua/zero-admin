package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponHistoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponHistoryUpdateLogic {
	return CouponHistoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponHistoryUpdateLogic) CouponHistoryUpdate(req types.UpdateCouponHistoryReq) (*types.UpdateCouponHistoryResp, error) {
	_, err := l.svcCtx.Sms.CouponHistoryUpdate(l.ctx, &smsclient.CouponHistoryUpdateReq{
		Id:             req.Id,
		CouponId:       req.CouponId,
		MemberId:       req.MemberId,
		CouponCode:     req.CouponCode,
		MemberNickname: req.MemberNickname,
		GetType:        req.GetType,
		CreateTime:     req.CreateTime,
		UseStatus:      req.UseStatus,
		UseTime:        req.UseTime,
		OrderId:        req.OrderId,
		OrderSn:        req.OrderSn,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateCouponHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
