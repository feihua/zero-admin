package orderreturnapplyservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderReturnApplyLogic 添加订单退货申请
/*
Author: LiuFeiHua
Date: 2024/6/12 10:14
*/
type AddOrderReturnApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderReturnApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderReturnApplyLogic {
	return &AddOrderReturnApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderReturnApply 添加订单退货申请
func (l *AddOrderReturnApplyLogic) AddOrderReturnApply(in *omsclient.AddOrderReturnApplyReq) (*omsclient.AddOrderReturnApplyResp, error) {
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

	return &omsclient.AddOrderReturnApplyResp{}, nil
}
