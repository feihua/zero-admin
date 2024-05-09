package couponproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *CouponProductCategoryRelationUpdateLogic) CouponProductCategoryRelationUpdate(in *smsclient.CouponProductCategoryRelationUpdateReq) (*smsclient.CouponProductCategoryRelationUpdateResp, error) {
	q := query.SmsCouponProductCategoryRelation
	_, err := q.WithContext(l.ctx).Updates(&model.SmsCouponProductCategoryRelation{
		ID:                  in.Id,
		CouponID:            in.CouponId,
		ProductCategoryID:   in.ProductCategoryId,
		ProductCategoryName: in.ProductCategoryName,
		ParentCategoryName:  in.ParentCategoryName,
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.CouponProductCategoryRelationUpdateResp{}, nil
}
