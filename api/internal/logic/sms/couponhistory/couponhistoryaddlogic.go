package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	_, err := l.svcCtx.Sms.CouponHistoryAdd(l.ctx, &smsclient.CouponHistoryAddReq{})

	if err != nil {
		return nil, err
	}

	return &types.AddCouponHistoryResp{}, nil
}
