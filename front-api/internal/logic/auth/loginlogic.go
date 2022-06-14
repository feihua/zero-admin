package auth

import (
	"context"
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

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginAndRegisterResp, err error) {
	memberAddResp, _ := l.svcCtx.Ums.MemberLogin(l.ctx, &umsclient.MemberLoginReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    "req.Mobile",
	})

	userInfo := types.UserInfo{
		NickName:  memberAddResp.Nickname,
		AvatarURL: memberAddResp.Icon,
	}

	data := types.LoginAndRegisterData{
		UserInfo: userInfo,
		Token:    memberAddResp.Token,
	}
	return &types.LoginAndRegisterResp{
		Errno:  0,
		Data:   data,
		Errmsg: "",
	}, nil
}
