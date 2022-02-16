package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ReturnApplyListLogic) ReturnApplyList(req types.ListReturnApplyReq) (*types.ListReturnApplyResp, error) {
	resp, err := l.svcCtx.Oms.OrderReturnApplyList(l.ctx, &omsclient.OrderReturnApplyListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询退货申请列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询退货申请失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtReturnApplyData

	for _, item := range resp.List {
		list = append(list, &types.ListtReturnApplyData{
			Id:               item.Id,
			OrderId:          item.OrderId,
			CompanyAddressId: item.CompanyAddressId,
			ProductId:        item.ProductId,
			OrderSn:          item.OrderSn,
			CreateTime:       item.CreateTime,
			MemberUsername:   item.MemberUsername,
			ReturnAmount:     float64(item.ReturnAmount),
			ReturnName:       item.ReturnName,
			ReturnPhone:      item.ReturnPhone,
			Status:           item.Status,
			HandleTime:       item.HandleTime,
			ProductPic:       item.ProductPic,
			ProductName:      item.ProductName,
			ProductBrand:     item.ProductBrand,
			ProductAttr:      item.ProductAttr,
			ProductCount:     item.ProductCount,
			ProductPrice:     float64(item.ProductPrice),
			ProductRealPrice: float64(item.ProductRealPrice),
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
		Message:  "查询退货申请失败",
	}, nil
}
