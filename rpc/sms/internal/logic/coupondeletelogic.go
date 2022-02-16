package logic

import (
	"context"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponDeleteLogic {
	return &CouponDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponDeleteLogic) CouponDelete(in *sms.CouponDeleteReq) (*sms.CouponDeleteResp, error) {
	err := l.svcCtx.SmsCouponModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sms.CouponDeleteResp{}, nil
}
