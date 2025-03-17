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

// UpdateOrderSettingIsDefaultLogic 更新订单设置默认状态
/*
Author: LiuFeiHua
Date: 2024/6/15 12:47
*/
type UpdateOrderSettingIsDefaultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderSettingIsDefaultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderSettingIsDefaultLogic {
	return &UpdateOrderSettingIsDefaultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateOrderSettingIsDefault 更新订单设置默认状态
func (l *UpdateOrderSettingIsDefaultLogic) UpdateOrderSettingIsDefault(req *types.UpdateOrderSettingIsDefaultReq) (resp *types.UpdateOrderSettingStatusResp, err error) {
	_, err = l.svcCtx.OrderSettingService.UpdateOrderSettingIsDefault(l.ctx, &omsclient.UpdateOrderSettingIsDefaultReq{
		Id:        req.Id,
		IsDefault: req.IsDefault,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新订单设置默认状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新订单设置默认状态失败")
	}

	return &types.UpdateOrderSettingStatusResp{
		Code:    "000000",
		Message: "更新订单设置默认状态成功",
	}, nil
}
