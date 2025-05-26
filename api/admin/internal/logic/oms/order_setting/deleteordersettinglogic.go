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

// DeleteOrderSettingLogic 删除订单设置
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type DeleteOrderSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOrderSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderSettingLogic {
	return &DeleteOrderSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteOrderSetting 删除订单设置
func (l *DeleteOrderSettingLogic) DeleteOrderSetting(req *types.DeleteOrderSettingReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.OrderSettingService.DeleteOrderSetting(l.ctx, &omsclient.DeleteOrderSettingReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除订单设置失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除订单设置成功",
	}, nil
}
