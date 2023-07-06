package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/omsmodel"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *OrderReturnApplyAddLogic) OrderReturnApplyAdd(in *oms.OrderReturnApplyAddReq) (*oms.OrderReturnApplyAddResp, error) {
	_, err := l.svcCtx.OmsOrderReturnApplyModel.Insert(l.ctx, &omsmodel.OmsOrderReturnApply{
		OrderId:          in.OrderId,
		ProductId:        in.ProductId,
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
		ProductPrice:     float64(in.ProductPrice),
		ProductRealPrice: float64(in.ProductRealPrice),
		Reason:           in.Reason,
		Description:      in.Description,
		ProofPics:        in.ProofPics,
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderReturnApplyAddResp{}, nil
}
