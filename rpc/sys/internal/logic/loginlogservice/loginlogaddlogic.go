package loginlogservicelogic

import (
	"context"
	"time"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogAddLogic {
	return &LoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogAddLogic) LoginLogAdd(in *sysclient.LoginLogAddReq) (*sysclient.LoginLogAddResp, error) {
	_, err := l.svcCtx.LoginLogModel.Insert(l.ctx, &sysmodel.SysLoginLog{
		UserName:   in.UserName,
		Status:     in.Status,
		Ip:         in.Ip,
		CreateTime: time.Now(),
		CreateBy:   in.CreateBy,
	})

	if err != nil {
		return nil, err
	}

	return &sysclient.LoginLogAddResp{}, nil
}
