package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/sys/sysclient"
	"strings"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//根据用户名和密码登录
func (l *UserLoginLogic) UserLogin(req types.LoginReq, ip string) (*types.LoginResp, error) {

	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("用户名或密码不能为空")
	}

	resp, err := l.svcCtx.Sys.Login(l.ctx, &sysclient.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("登录失败")
	}

	//保存登录日志
	_, _ = l.svcCtx.Sys.LoginLogAdd(l.ctx, &sysclient.LoginLogAddReq{
		UserName: resp.UserName,
		Status:   "login",
		Ip:       ip,
		CreateBy: resp.UserName,
	})

	return &types.LoginResp{
		Code:             "000000",
		Message:          "登录成功",
		Status:           resp.Status,
		CurrentAuthority: resp.CurrentAuthority,
		Id:               resp.Id,
		UserName:         resp.UserName,
		AccessToken:      resp.AccessToken,
		AccessExpire:     resp.AccessExpire,
		RefreshAfter:     resp.RefreshAfter,
	}, nil
}
