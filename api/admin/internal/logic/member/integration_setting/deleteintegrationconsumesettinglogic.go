package integration_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteIntegrationConsumeSettingLogic 删除积分消费设置
/*
Author: LiuFeiHua
Date: 2024/5/23 9:37
*/
type DeleteIntegrationConsumeSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteIntegrationConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteIntegrationConsumeSettingLogic {
	return &DeleteIntegrationConsumeSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteIntegrationConsumeSetting 删除积分消费设置
func (l *DeleteIntegrationConsumeSettingLogic) DeleteIntegrationConsumeSetting(req *types.DeleteIntegrationConsumeSettingReq) (resp *types.DeleteIntegrationConsumeSettingResp, err error) {
	_, err = l.svcCtx.IntegrationConsumeSettingService.DeleteIntegrationConsumeSetting(l.ctx, &umsclient.DeleteIntegrationConsumeSettingReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除积分消费设置异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除积分消费设置失败")
	}
	return &types.DeleteIntegrationConsumeSettingResp{
		Code:    "000000",
		Message: "删除积分消费设置成功",
	}, nil
}
