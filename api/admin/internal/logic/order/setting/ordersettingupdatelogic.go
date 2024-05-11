package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderSettingUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderSettingUpdateLogic {
	return OrderSettingUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderSettingUpdateLogic) OrderSettingUpdate(req types.UpdateOrderSettingReq) (*types.UpdateOrderSettingResp, error) {
	_, err := l.svcCtx.OrderSettingService.OrderSettingUpdate(l.ctx, &omsclient.OrderSettingUpdateReq{
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
