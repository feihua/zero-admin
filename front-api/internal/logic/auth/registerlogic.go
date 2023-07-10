package auth

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register 会员注册
func (l *RegisterLogic) Register(req types.RegisterReq) (resp *types.LoginAndRegisterResp, err error) {

	//调用rpc方法注册会员
	memberAddResp, registerResp, err2, done := RegisterMemberRpc(req, err, l)
	if done {
		return registerResp, err2
	}

	//构建返回数据
	return buildRegisterMemberResp(memberAddResp)
}

//构建返回数据
func buildRegisterMemberResp(memberAddResp *umsclient.MemberAddResp) (*types.LoginAndRegisterResp, error) {
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
		Errmsg: "会员注册成功",
	}, nil
}

// RegisterMemberRpc 调用rpc方法注册会员
func RegisterMemberRpc(req types.RegisterReq, err error, l *RegisterLogic) (*umsclient.MemberAddResp, *types.LoginAndRegisterResp, error, bool) {
	memberAddResp, err := l.svcCtx.MemberService.MemberAdd(l.ctx, &umsclient.MemberAddReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Mobile,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("会员注册失败,参数: %s,响应：%s", reqStr, err.Error())
		return nil, &types.LoginAndRegisterResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil, true
	}
	return memberAddResp, nil, nil, false
}
