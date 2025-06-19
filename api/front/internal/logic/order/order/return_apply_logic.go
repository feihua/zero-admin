package order

import (
	"context"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ReturnApplyLogic 申请退货
/*
Author: LiuFeiHua
Date: 2025/6/20 10:59
*/
type ReturnApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnApplyLogic {
	return &ReturnApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReturnApply 申请退货
func (l *ReturnApplyLogic) ReturnApply(req *types.ReturnApplyReq) (resp *types.ReturnApplyResp, err error) {
	_, err = l.svcCtx.OrderReturnApplyService.AddOrderReturnApply(l.ctx, &omsclient.AddOrderReturnApplyReq{
		OrderId:          req.OrderId,
		ProductId:        req.ProductId,
		OrderSn:          req.OrderSn,
		MemberUsername:   req.MemberUsername,
		ReturnName:       req.ReturnName,
		ReturnPhone:      req.ReturnPhone,
		Status:           0,
		ProductPic:       req.ProductPic,
		ProductName:      req.ProductName,
		ProductBrand:     req.ProductBrand,
		ProductAttr:      req.ProductAttr,
		ProductCount:     req.ProductCount,
		ProductPrice:     req.ProductPrice,
		ProductRealPrice: req.ProductRealPrice,
		Reason:           req.Reason,
		Description:      req.Description,
		ProofPics:        req.ProofPics,
	})
	if err != nil {
		logc.Errorf(l.ctx, "申请退货失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.ReturnApplyResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
