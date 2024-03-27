package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ConfigDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:11
*/
type ConfigDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigDeleteLogic {
	return ConfigDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ConfigDelete 删除配置
func (l *ConfigDeleteLogic) ConfigDelete(req types.DeleteConfigReq) (*types.DeleteConfigResp, error) {
	_, err := l.svcCtx.ConfigService.ConfigDelete(l.ctx, &sysclient.ConfigDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("删除参数配置失败")
	}

	return &types.DeleteConfigResp{}, nil
}
