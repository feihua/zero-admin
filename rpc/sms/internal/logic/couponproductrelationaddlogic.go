package logic

import (
	"context"
	"go-zero-admin/rpc/model/smsmodel"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *CouponProductRelationAddLogic) CouponProductRelationAdd(in *sms.CouponProductRelationAddReq) (*sms.CouponProductRelationAddResp, error) {
	_, err := l.svcCtx.SmsCouponProductRelationModel.Insert(smsmodel.SmsCouponProductRelation{
		CouponId:    in.CouponId,
		ProductId:   in.ProductId,
		ProductName: in.ProductName,
		ProductSn:   in.ProductSn,
	})
	if err != nil {
		return nil, err
	}

	return &sms.CouponProductRelationAddResp{}, nil
}
