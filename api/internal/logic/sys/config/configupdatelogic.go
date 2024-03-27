package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ConfigUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:11
*/
type ConfigUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigUpdateLogic {
	return ConfigUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ConfigUpdate 更新配置
func (l *ConfigUpdateLogic) ConfigUpdate(req types.UpdateConfigReq) (*types.UpdateConfigResp, error) {
	_, err := l.svcCtx.ConfigService.ConfigUpdate(l.ctx, &sysclient.ConfigUpdateReq{
		Id:           req.Id,
		Value:        req.Value,
		Label:        req.Label,
		Type:         req.Type,
		Description:  req.Description,
		Sort:         req.Sort,
		Remarks:      req.Remarks,
		LastUpdateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		return nil, errorx.NewDefaultError("更新参数配置失败")
	}

	return &types.UpdateConfigResp{}, nil
}
