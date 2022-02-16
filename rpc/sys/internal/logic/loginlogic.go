package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

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

//根据用户名和密码登录
func (l *LoginLogic) Login(in *sys.LoginReq) (*sys.LoginResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByName(in.UserName)

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", in.UserName, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", in.UserName, err.Error())
		return nil, err
	}

	if userInfo.Password != in.Password {
		logx.WithContext(l.ctx).Errorf("用户密码不正确,参数:%s", in.Password)
		return nil, errors.New("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.JWT.AccessSecret, now, l.svcCtx.Config.JWT.AccessExpire, userInfo.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	resp := &sys.LoginResp{
		Status:           "ok",
		CurrentAuthority: "admin",
		Id:               userInfo.Id,
		UserName:         userInfo.Name,
		AccessToken:      jwtToken,
		AccessExpire:     now + accessExpire,
		RefreshAfter:     now + accessExpire/2,
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(resp)
	logx.WithContext(l.ctx).Infof("登录成功,参数:%s,响应:%s", reqStr, listStr)
	return resp, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
