package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &types.AddIntegrationConsumeSettingResp{}, nil
}
