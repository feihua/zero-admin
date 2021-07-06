package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/sys/sysclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserDeleteLogic {
	return UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req types.DeleteUserReq) (*types.DeleteUserResp, error) {

	_, err := l.svcCtx.Sys.UserDelete(l.ctx, &sysclient.UserDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("删除用户失败")
	}

	return &types.DeleteUserResp{
		Code:    "000000",
		Message: "删除用户成功",
	}, nil
}
