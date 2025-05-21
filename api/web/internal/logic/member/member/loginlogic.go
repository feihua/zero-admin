package member

import (
	"context"
	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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
func (l *LoginLogic) Login(req *types.LoginReq, ip string) (resp *types.LoginResp, err1 error) {
	// 根据用户名,密码和手机号调用rpc会员登录方法
	result, err := l.svcCtx.MemberService.Login(l.ctx, &umsclient.LoginReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		Ip:       ip,
		Source:   req.Source,
	})

	if err != nil {
		logc.Errorf(l.ctx, "会员登录失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.LoginResp{
		Code:    0,
		Message: "登录成功",
		Data: types.LoginData{
			Token:     result.Token,
			TokenHead: "Bearer",
		},
	}, nil
}
