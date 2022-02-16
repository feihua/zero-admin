package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

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

func (l *LoginLogAddLogic) LoginLogAdd(in *sys.LoginLogAddReq) (*sys.LoginLogAddResp, error) {
	_, err := l.svcCtx.LoginLogModel.Insert(sysmodel.SysLoginLog{
		UserName:       in.UserName,
		Status:         in.Status,
		Ip:             in.Ip,
		CreateTime:     time.Now(),
		CreateBy:       in.CreateBy,
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &sys.LoginLogAddResp{}, nil
}
