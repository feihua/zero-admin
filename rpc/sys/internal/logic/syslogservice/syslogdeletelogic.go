package syslogservicelogic

import (
	"context"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogDeleteLogic {
	return &SysLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysLogDeleteLogic) SysLogDelete(in *sysclient.SysLogDeleteReq) (*sysclient.SysLogDeleteResp, error) {
	err := l.svcCtx.SysLogModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sysclient.SysLogDeleteResp{}, nil
}
