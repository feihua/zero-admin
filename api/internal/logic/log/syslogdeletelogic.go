package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &types.DeleteSysLogResp{}, nil
}
