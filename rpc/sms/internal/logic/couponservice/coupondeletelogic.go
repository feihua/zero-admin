package couponservicelogic

import (
	"context"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/sms/internal/svc"
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

func (l *CouponDeleteLogic) CouponDelete(in *smsclient.CouponDeleteReq) (*smsclient.CouponDeleteResp, error) {
	err := l.svcCtx.SmsCouponModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &smsclient.CouponDeleteResp{}, nil
}
