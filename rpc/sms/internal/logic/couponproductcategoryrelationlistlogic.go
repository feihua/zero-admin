package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CouponProductCategoryRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductCategoryRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductCategoryRelationListLogic {
	return &CouponProductCategoryRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductCategoryRelationListLogic) CouponProductCategoryRelationList(in *sms.CouponProductCategoryRelationListReq) (*sms.CouponProductCategoryRelationListResp, error) {
	// todo: add your logic here and delete this line

	return &sms.CouponProductCategoryRelationListResp{}, nil
}
