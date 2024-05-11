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

type IntegrationConsumeSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIntegrationConsumeSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) IntegrationConsumeSettingDeleteLogic {
	return IntegrationConsumeSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IntegrationConsumeSettingDeleteLogic) IntegrationConsumeSettingDelete(req types.DeleteIntegrationConsumeSettingReq) (*types.DeleteIntegrationConsumeSettingResp, error) {
	_, err := l.svcCtx.IntegrationConsumeSettingService.IntegrationConsumeSettingDelete(l.ctx, &umsclient.IntegrationConsumeSettingDeleteReq{
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
