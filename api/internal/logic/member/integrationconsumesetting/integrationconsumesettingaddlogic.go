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

type IntegrationConsumeSettingAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationConsumeSettingAddLogic {
	return IntegrationConsumeSettingAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingAddLogic) IntegrationConsumeSettingAdd(req types.AddIntegrationConsumeSettingReq) (*types.AddIntegrationConsumeSettingResp, error) {
	_, err := l.svcCtx.Ums.IntegrationConsumeSettingAdd(l.ctx, &umsclient.IntegrationConsumeSettingAddReq{
		DeductionPerAmount: req.DeductionPerAmount,
		MaxPercentPerOrder: req.MaxPercentPerOrder,
		UseUnit:            req.UseUnit,
		CouponStatus:       req.CouponStatus,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加积分消费设置信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加积分消费设置失败")
	}

	return &types.AddIntegrationConsumeSettingResp{
		Code:    "000000",
		Message: "添加积分消费设置成功",
	}, nil
}
