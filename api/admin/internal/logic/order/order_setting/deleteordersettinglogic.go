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

// DeleteOrderSettingLogic 删除订单设置
/*
Author: LiuFeiHua
Date: 2024/6/15 12:46
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
func (l *DeleteOrderSettingLogic) DeleteOrderSetting(req *types.DeleteOrderSettingReq) (resp *types.DeleteOrderSettingResp, err error) {
	_, err = l.svcCtx.OrderSettingService.DeleteOrderSetting(l.ctx, &omsclient.DeleteOrderSettingReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除订单设置异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除订单设置失败")
	}

	return &types.DeleteOrderSettingResp{
		Code:    "000000",
		Message: "",
	}, nil
}
