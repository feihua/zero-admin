package couponproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponProductCategoryRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductCategoryRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductCategoryRelationDeleteLogic {
	return &CouponProductCategoryRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductCategoryRelationDeleteLogic) CouponProductCategoryRelationDelete(in *smsclient.CouponProductCategoryRelationDeleteReq) (*smsclient.CouponProductCategoryRelationDeleteResp, error) {
	err := l.svcCtx.SmsCouponProductCategoryRelationModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &smsclient.CouponProductCategoryRelationDeleteResp{}, nil
}
