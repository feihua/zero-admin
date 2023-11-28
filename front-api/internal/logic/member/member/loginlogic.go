package member

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogic 会员登录
/*
Author: LiuFeiHua
Date: 2023/11/28 14:11
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
func (l *LoginLogic) Login(req *types.LoginReq, ip string) (resp *types.LoginResp, err error) {
	//根据用户名,密码和手机号调用rpc会员登录方法
	loginResp, err := l.svcCtx.MemberService.MemberLogin(l.ctx, &umsclient.MemberLoginReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Mobile,
		Ip:       ip,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logc.Errorf(l.ctx, "会员登录失败,参数: %s,响应：%s", reqStr, err.Error())
		return &types.LoginResp{
			Code:    1,
			Message: "登录失败",
		}, nil
	}

	return &types.LoginResp{
		Code:    0,
		Message: "登录成功",
		Data: types.LoginData{
			Token:     loginResp.Token,
			TokenHead: "Bearer ",
		},
	}, nil
}
