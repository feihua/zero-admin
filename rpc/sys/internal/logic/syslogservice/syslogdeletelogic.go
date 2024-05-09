package syslogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SysLogDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:08
*/
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

// SysLogDelete 删除操作日志
func (l *SysLogDeleteLogic) SysLogDelete(in *sysclient.SysLogDeleteReq) (*sysclient.SysLogDeleteResp, error) {
	q := query.SysLog
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &sysclient.SysLogDeleteResp{}, nil
}
