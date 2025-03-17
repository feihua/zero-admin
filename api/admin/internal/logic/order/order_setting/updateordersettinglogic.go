package order_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderSettingLogic 更新订单设置
/*
Author: LiuFeiHua
Date: 2024/6/15 12:47
*/
type UpdateOrderSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderSettingLogic {
	return &UpdateOrderSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateOrderSetting 更新订单设置
func (l *UpdateOrderSettingLogic) UpdateOrderSetting(req *types.UpdateOrderSettingReq) (resp *types.UpdateOrderSettingResp, err error) {
	_, err = l.svcCtx.OrderSettingService.UpdateOrderSetting(l.ctx, &omsclient.UpdateOrderSettingReq{
		Id:                  req.Id,                  // 主键
		FlashOrderOvertime:  req.FlashOrderOvertime,  // 秒杀订单超时关闭时间(分)
		NormalOrderOvertime: req.NormalOrderOvertime, // 正常订单超时时间(分)
		ConfirmOvertime:     req.ConfirmOvertime,     // 发货后自动确认收货时间（天）
		FinishOvertime:      req.FinishOvertime,      // 自动完成交易时间，不能申请售后（天）
		Status:              req.Status,              // 状态：0->禁用；1->启用
		IsDefault:           req.IsDefault,           // 是否默认：0->否；1->是
		CommentOvertime:     req.CommentOvertime,     // 订单完成后自动好评时间（天）
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新订单设置信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新订单设置失败")
	}

	return &types.UpdateOrderSettingResp{
		Code:    "000000",
		Message: "更新订单设置成功",
	}, nil
}
