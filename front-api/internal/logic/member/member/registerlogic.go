package member

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// RegisterLogic
/*
Author: LiuFeiHua
Date: 2024/3/14 上午9:48
*/
type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register 会员注册
func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	if req.Password != req.ConfirmPassword {
		return &types.RegisterResp{
			Code:    1,
			Message: "两次密码不一致",
		}, nil
	}
	rpcResult, err := l.svcCtx.MemberService.MemberAdd(l.ctx, &umsclient.MemberAddReq{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Mobile,
	})

	if err != nil {
		logc.Errorf(l.ctx, "会员注册失败,参数: %+v,响应：%s", req, err.Error())
		return &types.RegisterResp{
			Code:    1,
			Message: "注册失败",
		}, nil
	}
	return &types.RegisterResp{
		Code:    0,
		Message: "注册成功",
		Data: types.LoginData{
			Token:     rpcResult.Token,
			TokenHead: "Bearer",
		},
	}, nil
}
