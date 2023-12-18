package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ConfigAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:10
*/
type ConfigAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigAddLogic {
	return ConfigAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ConfigAdd 添加配置
func (l *ConfigAddLogic) ConfigAdd(req types.AddConfigReq) (*types.AddConfigResp, error) {
	_, err := l.svcCtx.ConfigService.ConfigAdd(l.ctx, &sysclient.ConfigAddReq{
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Sort:        req.Sort,
		Remarks:     req.Remarks,
		CreateBy:    l.ctx.Value("userName").(string),
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加参数配置失败")
	}

	return &types.AddConfigResp{}, nil
}
