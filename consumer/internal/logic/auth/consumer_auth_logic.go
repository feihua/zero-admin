package auth

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/feihua/zero-admin/consumer/internal/svc"
	"github.com/feihua/zero-admin/consumer/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConsumerAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumerAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConsumerAuthLogic {
	return &ConsumerAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumerAuthLogic) ConsumerAuth(req *types.LoginReq) (resp *types.LoginResp, err error) {
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	secret := l.svcCtx.Config.Auth.AccessSecret
	token, err := createJwtToken(secret, "member.Nickname", "member.Mobile", accessExpire, "member.MemberID")

	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &types.LoginResp{
		Code:    "000000",
		Message: "登录成功",
		Data: types.LoginData{
			AccessToken: token,
		},
	}, nil
}
func createJwtToken(secretKey, name, mobile string, seconds int64, memberId string) (string, error) {
	iat := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["memberId"] = memberId
	claims["memberName"] = name
	claims["mobile"] = mobile
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
