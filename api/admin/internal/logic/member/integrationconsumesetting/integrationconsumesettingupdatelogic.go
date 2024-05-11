package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

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
	_, err := l.svcCtx.IntegrationConsumeSettingService.IntegrationConsumeSettingUpdate(l.ctx, &umsclient.IntegrationConsumeSettingUpdateReq{
		Id:                 req.Id,
		DeductionPerAmount: req.DeductionPerAmount,
		MaxPercentPerOrder: req.MaxPercentPerOrder,
		UseUnit:            req.UseUnit,
		CouponStatus:       req.CouponStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新积分消费设置信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新积分消费设置失败")
	}

	return &types.UpdateIntegrationConsumeSettingResp{
		Code:    "000000",
		Message: "更新积分消费设置失败",
	}, nil
}
