package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryDeleteLogic {
	return &CouponHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponHistoryDeleteLogic) CouponHistoryDelete(in *sms.CouponHistoryDeleteReq) (*sms.CouponHistoryDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &sms.CouponHistoryDeleteResp{}, nil
}
