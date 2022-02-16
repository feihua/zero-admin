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

type CouponAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponAddLogic {
	return CouponAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponAddLogic) CouponAdd(req types.AddCouponReq) (*types.AddCouponResp, error) {
	_, err := l.svcCtx.Sms.CouponAdd(l.ctx, &smsclient.CouponAddReq{
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
		UseCount:     req.UseCount,
		ReceiveCount: req.ReceiveCount,
		EnableTime:   req.EnableTime,
		Code:         req.Code,
		MemberLevel:  req.MemberLevel,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加优惠券信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加优惠券失败")
	}

	return &types.AddCouponResp{
		Code:    "000000",
		Message: "添加优惠券成功",
	}, nil
}
