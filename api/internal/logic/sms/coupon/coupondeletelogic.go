package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponDeleteLogic {
	return CouponDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponDeleteLogic) CouponDelete(req types.DeleteCouponReq) (*types.DeleteCouponResp, error) {
	_, _ = l.svcCtx.Sms.CouponDelete(l.ctx, &smsclient.CouponDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteCouponResp{
		Code:    "000000",
		Message: "",
	}, nil
}
