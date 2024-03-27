package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// SysLogDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:20
*/
type SysLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SysLogDeleteLogic {
	return SysLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SysLogDelete 删除操作日志
func (l *SysLogDeleteLogic) SysLogDelete(req types.DeleteSysLogReq) (*types.DeleteSysLogResp, error) {
	_, err := l.svcCtx.SysLogService.SysLogDelete(l.ctx, &sysclient.SysLogDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("删除操作日志失败")
	}

	return &types.DeleteSysLogResp{
		Code:    "000000",
		Message: "",
	}, nil
}
