package userservicelogic

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:08
*/
type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 根据用户名和密码登录
func (l *LoginLogic) Login(in *sysclient.LoginReq) (*sysclient.LoginResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.UserName)

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logc.Errorf(l.ctx, "用户不存在,参数:%s,异常:%s", in.UserName, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logc.Errorf(l.ctx, "用户登录失败,参数:%s,异常:%s", in.UserName, err.Error())
		return nil, err
	}

	if userInfo.Password != in.Password {
		logc.Errorf(l.ctx, "用户密码不正确,参数:%s", in.Password)
		return nil, errors.New("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret
	jwtToken, err := l.getJwtToken(accessSecret, now, accessExpire, userInfo.Id, userInfo.Name)

	if err != nil {
		logc.Errorf(l.ctx, "生成token失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	resp := &sysclient.LoginResp{
		Status:           "ok",
		CurrentAuthority: "admin",
		Id:               userInfo.Id,
		UserName:         userInfo.Name,
		AccessToken:      jwtToken,
		AccessExpire:     now + accessExpire,
		RefreshAfter:     now + accessExpire/2,
	}

	logc.Infof(l.ctx, "登录成功,参数:%+v,响应:%+v", in, resp)
	return resp, nil
}

//生成jwt的token
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64, userName string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	claims["userName"] = userName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
