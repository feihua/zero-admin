package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Oms.OrderSettingUpdate(l.ctx, &omsclient.OrderSettingUpdateReq{
		Id:                  req.Id,
		FlashOrderOvertime:  req.FlashOrderOvertime,
		NormalOrderOvertime: req.NormalOrderOvertime,
		ConfirmOvertime:     req.ConfirmOvertime,
		FinishOvertime:      req.FinishOvertime,
		CommentOvertime:     req.CommentOvertime,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新订单设置信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新订单设置失败")
	}

	return &types.UpdateOrderSettingResp{
		Code:    "000000",
		Message: "更新订单设置成功",
	}, nil
}
