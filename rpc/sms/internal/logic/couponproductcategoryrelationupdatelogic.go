package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponProductCategoryRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductCategoryRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductCategoryRelationUpdateLogic {
	return &CouponProductCategoryRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductCategoryRelationUpdateLogic) CouponProductCategoryRelationUpdate(in *sms.CouponProductCategoryRelationUpdateReq) (*sms.CouponProductCategoryRelationUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sms.CouponProductCategoryRelationUpdateResp{}, nil
}
