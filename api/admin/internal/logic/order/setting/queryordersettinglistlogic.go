package setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderSettingListLogic {
	return &QueryOrderSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOrderSettingListLogic) QueryOrderSettingList(req *types.QueryOrderSettingListReq) (resp *types.QueryOrderSettingListResp, err error) {
	result, err := l.svcCtx.OrderSettingService.QueryOrderSettingList(l.ctx, &omsclient.QueryOrderSettingListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		Status:   req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询订单设置列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询订单设置失败")
	}

	var list []*types.QueryOrderSettingListData

	for _, item := range result.List {
		list = append(list, &types.QueryOrderSettingListData{
			Id:                  item.Id,                  //
			FlashOrderOvertime:  item.FlashOrderOvertime,  // 秒杀订单超时关闭时间(分)
			NormalOrderOvertime: item.NormalOrderOvertime, // 正常订单超时时间(分)
			ConfirmOvertime:     item.ConfirmOvertime,     // 发货后自动确认收货时间（天）
			FinishOvertime:      item.FinishOvertime,      // 自动完成交易时间，不能申请售后（天）
			Status:              item.Status,              // 状态：0->禁用；1->启用
			IsDefault:           item.IsDefault,           // 是否默认：0->否；1->是
			CommentOvertime:     item.CommentOvertime,     // 订单完成后自动好评时间（天）
		})
	}

	return &types.QueryOrderSettingListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询订单设置成功",
	}, nil
}
