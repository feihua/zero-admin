package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnApplyAddLogic {
	return ReturnApplyAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyAddLogic) ReturnApplyAdd(req types.AddReturnApplyReq) (*types.AddReturnApplyResp, error) {
	_, err := l.svcCtx.OrderReturnApplyService.OrderReturnApplyAdd(l.ctx, &omsclient.OrderReturnApplyAddReq{
		OrderId:          req.OrderId,
		ProductId:        req.ProductId,
		OrderSn:          req.OrderSn,
		MemberUsername:   req.MemberUsername,
		ReturnName:       req.ReturnName,
		ReturnPhone:      req.ReturnPhone,
		Status:           req.Status,
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加退货申请信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加退货申请失败")
	}

	return &types.AddReturnApplyResp{
		Code:    "000000",
		Message: "添加退货申请成功",
	}, nil
}
