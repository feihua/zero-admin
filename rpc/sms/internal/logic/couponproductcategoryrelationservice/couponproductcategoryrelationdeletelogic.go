package couponproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponProductCategoryRelationDeleteLogic 优惠券与产品分类关糸
/*
Author: LiuFeiHua
Date: 2024/5/8 10:08
*/
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

// CouponProductCategoryRelationDelete 删除优惠券与产品分类关糸
func (l *CouponProductCategoryRelationDeleteLogic) CouponProductCategoryRelationDelete(in *smsclient.CouponProductCategoryRelationDeleteReq) (*smsclient.CouponProductCategoryRelationDeleteResp, error) {
	q := query.SmsCouponProductCategoryRelation
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.CouponProductCategoryRelationDeleteResp{}, nil
}
