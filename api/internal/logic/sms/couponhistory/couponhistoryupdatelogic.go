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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新优惠券使用记录信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新优惠券使用记录失败")
	}

	return &types.UpdateCouponHistoryResp{
		Code:    "000000",
		Message: "更新优惠券使用记录成功",
	}, nil
}
