package member

import (
	"context"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

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
	rpcResult, err := l.svcCtx.MemberService.AddMember(l.ctx, &umsclient.AddMemberReq{
		MemberName: req.Username,
		Password:   req.Password,
		Phone:      req.Mobile,
	})

	if err != nil {
		logc.Errorf(l.ctx, "会员注册失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
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
