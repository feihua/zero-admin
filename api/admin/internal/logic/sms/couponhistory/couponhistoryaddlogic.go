package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponHistoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponHistoryAddLogic {
	return CouponHistoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponHistoryAddLogic) CouponHistoryAdd(req types.AddCouponHistoryReq) (*types.AddCouponHistoryResp, error) {
	_, err := l.svcCtx.CouponHistoryService.CouponHistoryAdd(l.ctx, &smsclient.CouponHistoryAddReq{
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
		logc.Errorf(l.ctx, "添加优惠券使用信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加优惠券使用记录失败")
	}

	return &types.AddCouponHistoryResp{
		Code:    "000000",
		Message: "添加优惠券使用记录成功",
	}, nil
}
