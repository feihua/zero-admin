package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
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
		reqStr, _ := json.Marshal(req)
		logx.Errorf("更新退货申请参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新退货申请失败")
	}

	return &types.UpdateReturnApplyResp{
		Code:    "000000",
		Message: "更新退货申请成功",
	}, nil
}
