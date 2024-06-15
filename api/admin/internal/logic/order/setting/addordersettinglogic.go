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

// AddOrderSettingLogic 添加订单设置
/*
Author: LiuFeiHua
Date: 2024/6/15 12:45
*/
type AddOrderSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrderSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderSettingLogic {
	return &AddOrderSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddOrderSetting 添加订单设置
func (l *AddOrderSettingLogic) AddOrderSetting(req *types.AddOrderSettingReq) (resp *types.AddOrderSettingResp, err error) {
	_, err = l.svcCtx.OrderSettingService.AddOrderSetting(l.ctx, &omsclient.AddOrderSettingReq{
		FlashOrderOvertime:  req.FlashOrderOvertime,
		NormalOrderOvertime: req.NormalOrderOvertime,
		ConfirmOvertime:     req.ConfirmOvertime,
		FinishOvertime:      req.FinishOvertime,
		CommentOvertime:     req.CommentOvertime,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加订单设置信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加订单设置失败")
	}

	return &types.AddOrderSettingResp{
		Code:    "000000",
		Message: "添加订单设置成功",
	}, nil
}
