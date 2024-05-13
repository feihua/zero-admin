package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderSettingListLogic {
	return OrderSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderSettingListLogic) OrderSettingList(req types.ListOrderSettingReq) (*types.ListOrderSettingResp, error) {
	resp, err := l.svcCtx.OrderSettingService.OrderSettingList(l.ctx, &omsclient.OrderSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询订单设置列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询订单设置失败")
	}

	var list []*types.ListOrderSettingData

	for _, item := range resp.List {
		list = append(list, &types.ListOrderSettingData{
			Id:                  item.Id,
			FlashOrderOvertime:  item.FlashOrderOvertime,
			NormalOrderOvertime: item.NormalOrderOvertime,
			ConfirmOvertime:     item.ConfirmOvertime,
			FinishOvertime:      item.FinishOvertime,
			CommentOvertime:     item.CommentOvertime,
		})
	}

	return &types.ListOrderSettingResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询订单设置成功",
	}, nil
}
