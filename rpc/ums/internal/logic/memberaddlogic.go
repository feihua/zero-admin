package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddLogic {
	return &MemberAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberAddLogic) MemberAdd(in *ums.MemberAddReq) (*ums.MemberAddResp, error) {

	member, _ := l.svcCtx.UmsMemberModel.FindOneByUsername(in.Username)
	if member != nil {
		logx.WithContext(l.ctx).Errorf("用户名已注册,参数:%s", in.Username)
		return nil, errors.New("用户名已注册")
	}

	phone, _ := l.svcCtx.UmsMemberModel.FindOneByPhone(in.Phone)
	if phone != nil {
		logx.WithContext(l.ctx).Errorf("手机号已注册,参数:%s", in.Phone)
		return nil, errors.New("手机号已注册")
	}

	result, _ := l.svcCtx.UmsMemberModel.Insert(umsmodel.UmsMember{
		MemberLevelId: 1,
		Username:      in.Username,
		Password:      in.Password,
		Nickname:      in.Username,
		Phone:         in.Phone,
		Status:        0,
		CreateTime:    time.Now(),
		Icon:          "",
	})

	userId, _ := result.LastInsertId()

	now := time.Now().Unix()
	//accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.JWT.AccessSecret, now, l.svcCtx.Config.JWT.AccessExpire, userId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	resp := &ums.MemberAddResp{
		Nickname: in.Username,
		Token:    jwtToken,
		Icon:     "",
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(resp)
	logx.WithContext(l.ctx).Infof("注册成功,参数:%s,响应:%s", reqStr, listStr)

	return resp, nil
}

func (l *MemberAddLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
