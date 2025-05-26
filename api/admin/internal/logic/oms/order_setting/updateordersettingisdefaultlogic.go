package order_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
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
func (l *UpdateOrderSettingIsDefaultLogic) UpdateOrderSettingIsDefault(req *types.UpdateOrderSettingStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.OrderSettingService.UpdateOrderSettingDefaultStatus(l.ctx, &omsclient.UpdateOrderSettingStatusReq{
		Id:       req.Id,
		Status:   req.Status,
		UpdateBy: userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新订单设置默认状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新订单设置默认状态失败")
	}

	return res.Success()
}
