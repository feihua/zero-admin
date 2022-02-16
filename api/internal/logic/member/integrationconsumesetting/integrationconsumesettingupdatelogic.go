package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IntegrationConsumeSettingUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationConsumeSettingUpdateLogic {
	return IntegrationConsumeSettingUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingUpdateLogic) IntegrationConsumeSettingUpdate(req types.UpdateIntegrationConsumeSettingReq) (*types.UpdateIntegrationConsumeSettingResp, error) {
	_, err := l.svcCtx.Ums.IntegrationConsumeSettingUpdate(l.ctx, &umsclient.IntegrationConsumeSettingUpdateReq{
		Id:                 req.Id,
		DeductionPerAmount: req.DeductionPerAmount,
		MaxPercentPerOrder: req.MaxPercentPerOrder,
		UseUnit:            req.UseUnit,
		CouponStatus:       req.CouponStatus,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新积分消费设置信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新积分消费设置失败")
	}

	return &types.UpdateIntegrationConsumeSettingResp{
		Code:    "000000",
		Message: "更新积分消费设置失败",
	}, nil
}
