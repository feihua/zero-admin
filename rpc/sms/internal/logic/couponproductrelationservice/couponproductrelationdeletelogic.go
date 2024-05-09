package couponproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponProductRelationDeleteLogic 优惠券与产品关糸
/*
Author: LiuFeiHua
Date: 2024/5/8 10:10
*/
type CouponProductRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductRelationDeleteLogic {
	return &CouponProductRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponProductRelationDelete 删除优惠券与产品关糸
func (l *CouponProductRelationDeleteLogic) CouponProductRelationDelete(in *smsclient.CouponProductRelationDeleteReq) (*smsclient.CouponProductRelationDeleteResp, error) {
	q := query.SmsCouponProductRelation
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.CouponProductRelationDeleteResp{}, nil
}
