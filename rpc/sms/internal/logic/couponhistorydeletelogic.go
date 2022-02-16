package logic

import (
	"context"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
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
	err := l.svcCtx.SmsCouponHistoryModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.CouponHistoryDeleteResp{}, nil
}
