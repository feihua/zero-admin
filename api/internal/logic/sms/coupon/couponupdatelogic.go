package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	_, err := l.svcCtx.Sms.CouponUpdate(l.ctx, &smsclient.CouponUpdateReq{})

	if err != nil {
		return nil, err
	}

	return &types.UpdateCouponResp{}, nil
}
