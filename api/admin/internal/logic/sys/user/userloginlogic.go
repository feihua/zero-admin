package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// UserLoginLogic 根据用户名(手机号)和密码登录
/*
Author: LiuFeiHua
Date: 2023/12/18 14:07
*/
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

// UserLogin 根据用户名(手机号)和密码登录
func (l *UserLoginLogic) UserLogin(req *types.LoginReq, ip, browser, os string) (*types.LoginResp, error) {

	resp, err := l.svcCtx.UserService.Login(l.ctx, &sysclient.LoginReq{
		Account:   strings.TrimSpace(req.Account),  // 用户名
		Password:  strings.TrimSpace(req.Password), // 密码
		IpAddress: ip,                              // ip地址
		Browser:   browser,                         // 浏览器
		Os:        os,                              // 操作系统
	})

	if err != nil {
		logc.Errorf(l.ctx, "用户登录：%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 把能访问的url存在在redis，在middleware中校验
	filed := strconv.FormatInt(resp.Id, 10)
	key := "zero:mall:token"
	err = l.svcCtx.Redis.HsetCtx(l.ctx, key, filed, strings.Join(resp.ApiUrls, ","))

	if err != nil {
		logc.Errorf(l.ctx, "设置用户：%s,权限到redis异常: %+v", resp.UserName, err)
	}
	return &types.LoginResp{
		Code:    "000000",
		Message: "登录成功",
		Data: types.LoginData{
			AccessToken: resp.AccessToken,
		},
	}, nil
}
