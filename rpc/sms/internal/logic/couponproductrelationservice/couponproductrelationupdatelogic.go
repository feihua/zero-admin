package couponproductrelationservicelogic

import (
	"context"
	"zero-admin/rpc/model/smsmodel"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/rpc/sms/internal/svc"

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
	err := l.svcCtx.SmsCouponProductRelationModel.Update(l.ctx, &smsmodel.SmsCouponProductRelation{
		Id:          in.Id,
		CouponId:    in.CouponId,
		ProductId:   in.ProductId,
		ProductName: in.ProductName,
		ProductSn:   in.ProductSn,
	})
	if err != nil {
		return nil, err
	}
	return &smsclient.CouponProductRelationUpdateResp{}, nil
}
