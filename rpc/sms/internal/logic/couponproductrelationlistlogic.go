package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductRelationListLogic {
	return &CouponProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductRelationListLogic) CouponProductRelationList(in *sms.CouponProductRelationListReq) (*sms.CouponProductRelationListResp, error) {
	// todo: add your logic here and delete this line

	return &sms.CouponProductRelationListResp{}, nil
}
