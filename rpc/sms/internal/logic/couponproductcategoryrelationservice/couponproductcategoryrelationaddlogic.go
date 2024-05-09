package couponproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponProductCategoryRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductCategoryRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductCategoryRelationAddLogic {
	return &CouponProductCategoryRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductCategoryRelationAddLogic) CouponProductCategoryRelationAdd(in *smsclient.CouponProductCategoryRelationAddReq) (*smsclient.CouponProductCategoryRelationAddResp, error) {
	err := query.SmsCouponProductCategoryRelation.WithContext(l.ctx).Create(&model.SmsCouponProductCategoryRelation{
		CouponID:            in.CouponId,
		ProductCategoryID:   in.ProductCategoryId,
		ProductCategoryName: in.ProductCategoryName,
		ParentCategoryName:  in.ParentCategoryName,
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.CouponProductCategoryRelationAddResp{}, nil
}
