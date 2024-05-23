package log

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOperateLogLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:20
*/
type DeleteOperateLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOperateLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteOperateLogLogic {
	return DeleteOperateLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteOperateLog 删除操作日志
func (l *DeleteOperateLogLogic) DeleteOperateLog(req *types.DeleteSysLogReq) (*types.DeleteSysLogResp, error) {
	_, err := l.svcCtx.SysLogService.SysLogDelete(l.ctx, &sysclient.SysLogDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("删除操作日志失败")
	}

	return &types.DeleteSysLogResp{
		Code:    "000000",
		Message: "删除操作日志成功",
	}, nil
}
