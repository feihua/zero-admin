package logic

import (
	"context"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

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

func (l *ConfigDeleteLogic) ConfigDelete(in *sys.ConfigDeleteReq) (*sys.ConfigDeleteResp, error) {
	err := l.svcCtx.DeptModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sys.ConfigDeleteResp{}, nil
}
