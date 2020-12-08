package logic

import (
	"context"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnApplyUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnApplyUpdateLogic {
	return ReturnApplyUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyUpdateLogic) ReturnApplyUpdate(req types.UpdateReturnApplyReq) (*types.UpdateReturnApplyResp, error) {
	_, err := l.svcCtx.Oms.OrderReturnApplyUpdate(l.ctx, &omsclient.OrderReturnApplyUpdateReq{
		Id:               req.Id,
		OrderId:          req.OrderId,
		CompanyAddressId: req.CompanyAddressId,
		ProductId:        req.ProductId,
		OrderSn:          req.OrderSn,
		CreateTime:       req.CreateTime,
		MemberUsername:   req.MemberUsername,
		ReturnAmount:     int64(req.ReturnAmount),
		ReturnName:       req.ReturnName,
		ReturnPhone:      req.ReturnPhone,
		Status:           req.Status,
		HandleTime:       req.HandleTime,
		ProductPic:       req.ProductPic,
		ProductName:      req.ProductName,
		ProductBrand:     req.ProductBrand,
		ProductAttr:      req.ProductAttr,
		ProductCount:     req.ProductCount,
		ProductPrice:     int64(req.ProductPrice),
		ProductRealPrice: int64(req.ProductRealPrice),
		Reason:           req.Reason,
		Description:      req.Description,
		ProofPics:        req.ProofPics,
		HandleNote:       req.HandleNote,
		HandleMan:        req.HandleMan,
		ReceiveMan:       req.ReceiveMan,
		ReceiveTime:      req.ReceiveTime,
		ReceiveNote:      req.ReceiveNote,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateReturnApplyResp{}, nil
}
