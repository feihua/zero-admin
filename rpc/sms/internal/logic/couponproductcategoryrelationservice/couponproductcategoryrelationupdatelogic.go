package couponproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/smsmodel"
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
	err := l.svcCtx.SmsCouponProductCategoryRelationModel.Update(l.ctx, &smsmodel.SmsCouponProductCategoryRelation{
		Id:                  in.Id,
		CouponId:            in.CouponId,
		ProductCategoryId:   in.ProductCategoryId,
		ProductCategoryName: in.ProductCategoryName,
		ParentCategoryName:  in.ParentCategoryName,
	})
	if err != nil {
		return nil, err
	}

	return &smsclient.CouponProductCategoryRelationUpdateResp{}, nil
}
