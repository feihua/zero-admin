package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &types.UpdateCouponHistoryResp{}, nil
}
