package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponUpdateLogic {
	return CouponUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponUpdateLogic) CouponUpdate(req types.UpdateCouponReq) (*types.UpdateCouponResp, error) {
	_, err := l.svcCtx.CouponService.CouponUpdate(l.ctx, &smsclient.CouponUpdateReq{
		Id:           req.Id,
		Type:         req.Type,
		Name:         req.Name,
		Platform:     req.Platform,
		Count:        req.Count,
		Amount:       req.Amount,
		PerLimit:     req.PerLimit,
		MinPoint:     req.MinPoint,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		UseType:      req.UseType,
		Note:         req.Note,
		PublishCount: req.PublishCount,
		EnableTime:   req.EnableTime,
		Code:         req.Code,
		MemberLevel:  req.MemberLevel,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新优惠券信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新优惠券失败")
	}

	return &types.UpdateCouponResp{
		Code:    "000000",
		Message: "更新优惠券成功",
	}, nil
}
