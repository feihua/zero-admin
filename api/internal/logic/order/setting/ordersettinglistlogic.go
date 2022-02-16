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
	resp, err := l.svcCtx.Oms.OrderSettingList(l.ctx, &omsclient.OrderSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询订单设置列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询订单设置失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtOrderSettingData

	for _, item := range resp.List {
		list = append(list, &types.ListtOrderSettingData{
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
