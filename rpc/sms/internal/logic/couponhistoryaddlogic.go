package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponHistoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryAddLogic {
	return &CouponHistoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponHistoryAddLogic) CouponHistoryAdd(in *sms.CouponHistoryAddReq) (*sms.CouponHistoryAddResp, error) {
	// todo: add your logic here and delete this line

	return &sms.CouponHistoryAddResp{}, nil
}
