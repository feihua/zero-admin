package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Sms.CouponHistoryAdd(l.ctx, &smsclient.CouponHistoryAddReq{
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加优惠券使用信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加优惠券使用记录失败")
	}

	return &types.AddCouponHistoryResp{
		Code:    "000000",
		Message: "添加优惠券使用记录成功",
	}, nil
}
