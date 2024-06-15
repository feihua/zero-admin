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
		Id:                  req.Id,
		FlashOrderOvertime:  req.FlashOrderOvertime,
		NormalOrderOvertime: req.NormalOrderOvertime,
		ConfirmOvertime:     req.ConfirmOvertime,
		FinishOvertime:      req.FinishOvertime,
		CommentOvertime:     req.CommentOvertime,
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
