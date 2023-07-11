package member

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *LoginLogic) Login(req *types.LoginReq, ip string) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.MemberService.MemberLogin(l.ctx, &umsclient.MemberLoginReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Mobile,
		Ip:       ip,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("会员登录失败,参数: %s,响应：%s", reqStr, err.Error())
		return &types.LoginResp{
			Code:    1,
			Message: "登录失败",
		}, nil
	}

	return &types.LoginResp{
		Code:    0,
		Message: "登录成功",
		Data:    loginResp.Token,
	}, nil
}
