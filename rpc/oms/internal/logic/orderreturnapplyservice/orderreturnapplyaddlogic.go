package orderreturnapplyservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderReturnApplyAddLogic 申请退货
/*
Author: LiuFeiHua
Date: 2024/5/16 14:33
*/
type OrderReturnApplyAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnApplyAddLogic {
	return &OrderReturnApplyAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderReturnApplyAdd 申请退货
func (l *OrderReturnApplyAddLogic) OrderReturnApplyAdd(in *omsclient.OrderReturnApplyAddReq) (*omsclient.OrderReturnApplyAddResp, error) {
	err := query.OmsOrderReturnApply.WithContext(l.ctx).Create(&model.OmsOrderReturnApply{
		OrderID:          in.OrderId,
		ProductID:        in.ProductId,
		OrderSn:          in.OrderSn,
		CreateTime:       time.Now(),
		MemberUsername:   in.MemberUsername,
		ReturnName:       in.ReturnName,
		ReturnPhone:      in.ReturnPhone,
		Status:           in.Status,
		ProductPic:       in.ProductPic,
		ProductName:      in.ProductName,
		ProductBrand:     in.ProductBrand,
		ProductAttr:      in.ProductAttr,
		ProductCount:     in.ProductCount,
		ProductPrice:     in.ProductPrice,
		ProductRealPrice: in.ProductRealPrice,
		Reason:           in.Reason,
		Description:      in.Description,
		ProofPics:        in.ProofPics,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.OrderReturnApplyAddResp{}, nil
}
