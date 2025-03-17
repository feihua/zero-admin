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

// UpdateOrderSettingStatusLogic 更新订单设置状态
type UpdateOrderSettingStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderSettingStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderSettingStatusLogic {
	return &UpdateOrderSettingStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateOrderSettingStatus 更新订单设置状态
func (l *UpdateOrderSettingStatusLogic) UpdateOrderSettingStatus(req *types.UpdateOrderSettingStatusReq) (resp *types.UpdateOrderSettingStatusResp, err error) {
	_, err = l.svcCtx.OrderSettingService.UpdateOrderSettingStatus(l.ctx, &omsclient.UpdateOrderSettingStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新订单设置状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新订单设置状态失败")
	}

	return &types.UpdateOrderSettingStatusResp{
		Code:    "000000",
		Message: "更新订单设置状态成功",
	}, nil
}
