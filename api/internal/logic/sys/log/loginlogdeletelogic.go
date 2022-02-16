package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogDeleteLogic {
	return LoginLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogDeleteLogic) LoginLogDelete(req types.DeleteLoginLogReq) (*types.DeleteLoginLogResp, error) {
	_, err := l.svcCtx.Sys.LoginLogDelete(l.ctx, &sysclient.LoginLogDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("删除登录日志失败")
	}

	return &types.DeleteLoginLogResp{
		Code:    "000000",
		Message: "删除登录日志成功",
	}, nil
}
