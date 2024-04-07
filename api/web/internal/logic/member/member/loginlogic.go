package member

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogic
/*
Author: LiuFeiHua
Date: 2024/4/7 18:02
*/
type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 会员登录
func (l *LoginLogic) Login(req *types.LoginReq, ip string) (*types.LoginResp, error) {
	//根据用户名,密码和手机号调用rpc会员登录方法
	rpcResult, err := l.svcCtx.MemberService.MemberLogin(l.ctx, &umsclient.MemberLoginReq{
		Account:  req.Account,
		Password: req.Password,
		Ip:       ip,
	})

	if err != nil {
		logc.Errorf(l.ctx, "会员登录失败,参数: %+v,响应：%s", req, err.Error())
		return &types.LoginResp{
			Code:    1,
			Message: "登录失败",
		}, nil
	}

	return &types.LoginResp{
		Code:    0,
		Message: "登录成功",
		Data: types.LoginData{
			Token:     rpcResult.Token,
			TokenHead: "Bearer",
		},
	}, nil
}
