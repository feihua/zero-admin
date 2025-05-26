package order_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderSettingStatusLogic 更新订单设置状态状态
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
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
func (l *UpdateOrderSettingStatusLogic) UpdateOrderSettingStatus(req *types.UpdateOrderSettingStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.OrderSettingService.UpdateOrderSettingStatus(l.ctx, &omsclient.UpdateOrderSettingStatusReq{
		Id:       req.Id,     // 主键ID
		Status:   req.Status, // 状态：0->禁用；1->启用
		UpdateBy: userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新订单设置状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新订单设置状态成功",
	}, nil
}
