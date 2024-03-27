package loginlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:07
*/
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

// LoginLogDelete 删除登录日志
func (l *LoginLogDeleteLogic) LoginLogDelete(in *sysclient.LoginLogDeleteReq) (*sysclient.LoginLogDeleteResp, error) {
	err := l.svcCtx.LoginLogModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &sysclient.LoginLogDeleteResp{}, nil
}
