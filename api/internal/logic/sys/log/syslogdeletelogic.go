package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *SysLogDeleteLogic) SysLogDelete(req types.DeleteSysLogReq) (*types.DeleteSysLogResp, error) {
	_, err := l.svcCtx.Sys.SysLogDelete(l.ctx, &sysclient.SysLogDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("删除操作日志失败")
	}

	return &types.DeleteSysLogResp{
		Code:    "000000",
		Message: "",
	}, nil
}
