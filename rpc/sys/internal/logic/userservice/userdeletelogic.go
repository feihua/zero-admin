package userservicelogic

import (
	"context"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteLogic) UserDelete(in *sysclient.UserDeleteReq) (*sysclient.UserDeleteResp, error) {
	_ = l.svcCtx.UserModel.DeleteByIds(l.ctx, in.Ids)

	return &sysclient.UserDeleteResp{}, nil
}
