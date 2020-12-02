package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponHistoryDeleteLogic {
	return CouponHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponHistoryDeleteLogic) CouponHistoryDelete(req types.DeleteCouponHistoryReq) (*types.DeleteCouponHistoryResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteCouponHistoryResp{}, nil
}
