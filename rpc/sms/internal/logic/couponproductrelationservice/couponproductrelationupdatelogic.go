package couponproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponProductRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductRelationUpdateLogic {
	return &CouponProductRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductRelationUpdateLogic) CouponProductRelationUpdate(in *smsclient.CouponProductRelationUpdateReq) (*smsclient.CouponProductRelationUpdateResp, error) {
	q := query.SmsCouponProductRelation
	_, err := q.WithContext(l.ctx).Updates(&model.SmsCouponProductRelation{
		ID:          in.Id,
		CouponID:    in.CouponId,
		ProductID:   in.ProductId,
		ProductName: in.ProductName,
		ProductSn:   in.ProductSn,
	})

	if err != nil {
		return nil, err
	}
	return &smsclient.CouponProductRelationUpdateResp{}, nil
}
