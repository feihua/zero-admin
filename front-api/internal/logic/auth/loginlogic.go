package auth

import (
	"context"
	"zero-admin/rpc/ums/ums"

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

// Login 会员登录
func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginAndRegisterResp, err error) {

	memberLoginResp, loginResp, err2, done := callLoginRpc(req, err, l)
	if done {
		return loginResp, err2
	}

	return buildLoginResp(memberLoginResp)
}

// 构建返回数据
func buildLoginResp(memberAddResp *ums.MemberLoginResp) (*types.LoginAndRegisterResp, error) {
	userInfo := types.UserInfo{
		NickName: memberAddResp.Nickname,
		Icon:     memberAddResp.Icon,
	}

	data := types.LoginAndRegisterData{
		UserInfo: userInfo,
		Token:    memberAddResp.Token,
	}
	return &types.LoginAndRegisterResp{
		Errno:  0,
		Data:   data,
		Errmsg: "会员登录成功",
	}, nil
}

// 调用rpc方法登录
func callLoginRpc(req types.LoginReq, err error, l *LoginLogic) (*ums.MemberLoginResp, *types.LoginAndRegisterResp, error, bool) {
	// loginResp, err = l.svcCtx.Ums.MemberLogin(l.ctx, &umsclient.MemberLoginReq{
	// 	Username: req.Username,
	// 	Password: req.Password,
	// 	Phone:    "req.Mobile",
	// })

	// if err != nil {
	// 	reqStr, _ := json.Marshal(req)
	// 	logx.WithContext(l.ctx).Errorf("会员登录失败,参数: %s,响应：%s", reqStr, err.Error())
	// 	return nil, &types.LoginAndRegisterResp{
	// 		Errno:  1,
	// 		Errmsg: err.Error(),
	// 	}, nil, true
	// }
	return nil, nil, nil, false
}
