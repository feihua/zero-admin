package couponproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponProductRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductRelationAddLogic {
	return &CouponProductRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductRelationAddLogic) CouponProductRelationAdd(in *smsclient.CouponProductRelationAddReq) (*smsclient.CouponProductRelationAddResp, error) {
	err := query.SmsCouponProductRelation.WithContext(l.ctx).Create(&model.SmsCouponProductRelation{
		CouponID:    in.CouponId,
		ProductID:   in.ProductId,
		ProductName: in.ProductName,
		ProductSn:   in.ProductSn,
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.CouponProductRelationAddResp{}, nil
}
