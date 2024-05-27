package loginlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogAddLogic 添加登录日志
/*
Author: LiuFeiHua
Date: 2023/12/18 17:07
*/
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

// LoginLogAdd 添加登录日志
func (l *LoginLogAddLogic) LoginLogAdd(in *sysclient.LoginLogAddReq) (*sysclient.LoginLogAddResp, error) {
	loginLog := &model.SysLoginLog{
		LoginName:   in.UserName,
		LoginStatus: in.Status,
		LoginIP:     in.Ip,
		LoginTime:   time.Now(),
	}
	err := query.SysLoginLog.WithContext(l.ctx).Create(loginLog)
	if err != nil {
		logc.Errorf(l.ctx, "添加登录日志失败,参数:%+v,异常:%s", loginLog, err.Error())
		return nil, errors.New("添加登录日志失败")
	}

	return &sysclient.LoginLogAddResp{}, nil
}
