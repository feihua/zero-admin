package logic

import (
	"context"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogDeleteLogic {
	return &LoginLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogDeleteLogic) LoginLogDelete(in *sys.LoginLogDeleteReq) (*sys.LoginLogDeleteResp, error) {
	err := l.svcCtx.LoginLogModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &sys.LoginLogDeleteResp{}, nil
}
