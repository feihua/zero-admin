package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ReturnApplyListLogic 退货申请
/*
Author: LiuFeiHua
Date: 2024/5/14 14:45
*/
type ReturnApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnApplyListLogic {
	return ReturnApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReturnApplyList 查询退货申请
func (l *ReturnApplyListLogic) ReturnApplyList(req types.ListReturnApplyReq) (*types.ListReturnApplyResp, error) {
	resp, err := l.svcCtx.OrderReturnApplyService.OrderReturnApplyList(l.ctx, &omsclient.OrderReturnApplyListReq{
		Current:        req.Current,
		PageSize:       req.PageSize,
		OrderSn:        strings.TrimSpace(req.OrderSn),
		MemberUsername: strings.TrimSpace(req.MemberUsername),
		HandleTime:     req.HandleTime,
		CreateTime:     req.CreateTime,
		Status:         req.Status,
		HandleMan:      strings.TrimSpace(req.HandleMan),
		ReturnName:     strings.TrimSpace(req.ReturnName),
		ReturnPhone:    strings.TrimSpace(req.ReturnPhone),
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询退货申请列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询退货申请列表失败")
	}

	var list []*types.ListReturnApplyData

	for _, item := range resp.List {
		list = append(list, &types.ListReturnApplyData{
			Id:               item.Id,
			OrderId:          item.OrderId,
			CompanyAddressId: item.CompanyAddressId,
			ProductId:        item.ProductId,
			OrderSn:          item.OrderSn,
			CreateTime:       item.CreateTime,
			MemberUsername:   item.MemberUsername,
			ReturnAmount:     item.ReturnAmount,
			ReturnName:       item.ReturnName,
			ReturnPhone:      item.ReturnPhone,
			Status:           item.Status,
			HandleTime:       item.HandleTime,
			ProductPic:       item.ProductPic,
			ProductName:      item.ProductName,
			ProductBrand:     item.ProductBrand,
			ProductAttr:      item.ProductAttr,
			ProductCount:     item.ProductCount,
			ProductPrice:     item.ProductPrice,
			ProductRealPrice: item.ProductRealPrice,
			Reason:           item.Reason,
			Description:      item.Description,
			ProofPics:        item.ProofPics,
			HandleNote:       item.HandleNote,
			HandleMan:        item.HandleMan,
			ReceiveMan:       item.ReceiveMan,
			ReceiveTime:      item.ReceiveTime,
			ReceiveNote:      item.ReceiveNote,
		})
	}

	return &types.ListReturnApplyResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询退货申请列表成功",
	}, nil
}
