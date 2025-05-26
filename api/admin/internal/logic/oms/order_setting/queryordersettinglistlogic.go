package order_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderSettingListLogic 查询订单设置列表
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
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

// QueryOrderSettingList 查询订单设置列表
func (l *QueryOrderSettingListLogic) QueryOrderSettingList(req *types.QueryOrderSettingListReq) (resp *types.QueryOrderSettingListResp, err error) {
	result, err := l.svcCtx.OrderSettingService.QueryOrderSettingList(l.ctx, &omsclient.QueryOrderSettingListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		Status:    req.Status,    // 状态：0->禁用；1->启用
		IsDefault: req.IsDefault, // 是否默认：0->否；1->是
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字订单设置列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryOrderSettingListData

	for _, detail := range result.List {
		list = append(list, &types.QueryOrderSettingListData{
			Id:                  detail.Id,                  // 主键ID
			FlashOrderOvertime:  detail.FlashOrderOvertime,  // 秒杀订单超时关闭时间(分)
			NormalOrderOvertime: detail.NormalOrderOvertime, // 正常订单超时时间(分)
			ConfirmOvertime:     detail.ConfirmOvertime,     // 发货后自动确认收货时间（天）
			FinishOvertime:      detail.FinishOvertime,      // 自动完成交易时间，不能申请售后（天）
			Status:              detail.Status,              // 状态：0->禁用；1->启用
			IsDefault:           detail.IsDefault,           // 是否默认：0->否；1->是
			CommentOvertime:     detail.CommentOvertime,     // 订单完成后自动好评时间（天）
			CreateBy:            detail.CreateBy,            // 创建人ID
			CreateTime:          detail.CreateTime,          // 创建时间
			UpdateBy:            detail.UpdateBy,            // 更新人ID
			UpdateTime:          detail.UpdateTime,          // 更新时间

		})
	}

	return &types.QueryOrderSettingListResp{
		Code:     "000000",
		Message:  "查询订单设置列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
