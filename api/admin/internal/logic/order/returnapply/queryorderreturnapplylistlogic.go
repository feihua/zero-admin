package returnapply

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

// QueryOrderReturnApplyListLogic 查询退货申请
/*
Author: LiuFeiHua
Date: 2024/6/15 11:38
*/
type QueryOrderReturnApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderReturnApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnApplyListLogic {
	return &QueryOrderReturnApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOrderReturnApplyList 查询退货申请
func (l *QueryOrderReturnApplyListLogic) QueryOrderReturnApplyList(req *types.QueryOrderReturnApplyListReq) (resp *types.QueryOrderReturnApplyListResp, err error) {
	result, err := l.svcCtx.OrderReturnApplyService.QueryOrderReturnApplyList(l.ctx, &omsclient.QueryOrderReturnApplyListReq{
		PageNum:        req.Current,
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

	var list []*types.QueryOrderReturnApplyListData

	for _, item := range result.List {
		list = append(list, &types.QueryOrderReturnApplyListData{
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

	return &types.QueryOrderReturnApplyListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询退货申请列表成功",
	}, nil
}
