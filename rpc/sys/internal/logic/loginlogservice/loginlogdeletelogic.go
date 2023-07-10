package loginlogservicelogic

import (
	"context"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

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

func (l *LoginLogDeleteLogic) LoginLogDelete(in *sysclient.LoginLogDeleteReq) (*sysclient.LoginLogDeleteResp, error) {
	err := l.svcCtx.LoginLogModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sysclient.LoginLogDeleteResp{}, nil
}
