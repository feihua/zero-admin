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
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	HandleTime, _ := time.Parse("2006-01-02 15:04:05", in.HandleTime)
	ReceiveTime, _ := time.Parse("2006-01-02 15:04:05", in.ReceiveTime)
	_, err := l.svcCtx.OmsOrderReturnApplyModel.Insert(omsmodel.OmsOrderReturnApply{
		OrderId:          in.OrderId,
		CompanyAddressId: in.CompanyAddressId,
		ProductId:        in.ProductId,
		OrderSn:          in.OrderSn,
		CreateTime:       CreateTime,
		MemberUsername:   in.MemberUsername,
		ReturnAmount:     float64(in.ReturnAmount),
		ReturnName:       in.ReturnName,
		ReturnPhone:      in.ReturnPhone,
		Status:           in.Status,
		HandleTime:       HandleTime,
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
		HandleNote:       in.HandleNote,
		HandleMan:        in.HandleMan,
		ReceiveMan:       in.ReceiveMan,
		ReceiveTime:      ReceiveTime,
	})
	if err != nil {
		return nil, err
	}

	return &oms.OrderReturnApplyAddResp{}, nil
}
