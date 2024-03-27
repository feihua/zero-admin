package couponproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/smsmodel"
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
	_, err := l.svcCtx.SmsCouponProductCategoryRelationModel.Insert(l.ctx, &smsmodel.SmsCouponProductCategoryRelation{
		CouponId:            in.CouponId,
		ProductCategoryId:   in.ProductCategoryId,
		ProductCategoryName: in.ProductCategoryName,
		ParentCategoryName:  in.ParentCategoryName,
	})
	if err != nil {
		return nil, err
	}

	return &smsclient.CouponProductCategoryRelationAddResp{}, nil
}
