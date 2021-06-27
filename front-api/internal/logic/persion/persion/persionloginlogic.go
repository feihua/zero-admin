package logic

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go-zero-admin/rpc/us/usclient"
	"strings"
	"time"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PersionLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPersionLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) PersionLoginLogic {
	return PersionLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PersionLoginLogic) PersionLogin(req types.PersionLoginReq) (*types.PersionLoginResp, error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.PhoneNumber)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil,errors.New("参数错误")
	}

	resp, err := l.svcCtx.Us.PersionLogin(l.ctx, &usclient.PersionLoginReq{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})

	if err != nil {
		return nil,err
	}

	//保存登录日志
	//_, _ = l.svcCtx.Sys.LoginLogAdd(l.ctx, &sysclient.LoginLogAddReq{
	//	UserName: resp.UserName,
	//	Status:   "login",
	//	Ip:       ip,
	//	CreateBy: resp.UserName,
	//})
	// ---start---
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, resp.Id)
	if err != nil {
		return nil, err
	}
	// ---end---

	return &types.PersionLoginResp{
		Code:    0,
		Message: "success",
		Data:    types.PersionLoginRespData{
			Id:           resp.Id,
			Name:         resp.Name,
			PhoneNumber:  resp.PhoneNumber,
			Email:        resp.Email,
			RoleId:       resp.RoleId,
			RoleName:     resp.RoleName,
			Gender:       resp.Gender,
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	},nil
}

func (l *PersionLoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
