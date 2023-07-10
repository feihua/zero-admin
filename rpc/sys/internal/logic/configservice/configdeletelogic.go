package configservicelogic

import (
	"context"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigDeleteLogic {
	return &ConfigDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigDeleteLogic) ConfigDelete(in *sysclient.ConfigDeleteReq) (*sysclient.ConfigDeleteResp, error) {
	err := l.svcCtx.DeptModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sysclient.ConfigDeleteResp{}, nil
}
