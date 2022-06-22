package logic

import (
	"context"
	"database/sql"
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

// MemberAdd 会员注册
func (l *MemberAddLogic) MemberAdd(in *ums.MemberAddReq) (*ums.MemberAddResp, error) {

	//1.校验参数
	err := checkParams(in, l)
	if err != nil {
		return nil, err
	}

	//2.插入数据
	result, _ := l.svcCtx.UmsMemberModel.Insert(umsmodel.UmsMember{
		MemberLevelId:         4, //默认是普通会员
		Username:              in.Username,
		Password:              in.Password,
		Nickname:              in.Username,
		Phone:                 in.Phone,
		Status:                0,
		Icon:                  "",
		Gender:                0,
		Birthday:              time.Now(),
		City:                  "",
		Job:                   "",
		PersonalizedSignature: "",
		SourceType:            0,
		Integration:           0,
		Growth:                0,
		LuckeyCount:           0,
		HistoryIntegration:    0,
	})

	//3.构建返回数据
	return buildMemberRegisterResp(in, result, l)
}

//校验参数
func checkParams(in *ums.MemberAddReq, l *MemberAddLogic) error {
	member, _ := l.svcCtx.UmsMemberModel.FindOneByUsername(in.Username)
	if member != nil {
		logx.WithContext(l.ctx).Errorf("用户名已注册,参数Username:%s", in.Username)
		return errors.New("用户名已注册")
	}

	phone, _ := l.svcCtx.UmsMemberModel.FindOneByPhone(in.Phone)
	if phone != nil {
		logx.WithContext(l.ctx).Errorf("手机号已注册,参数Phone:%s", in.Phone)
		return errors.New("手机号已注册")
	}
	return nil
}

//构建返回数据
func buildMemberRegisterResp(in *ums.MemberAddReq, result sql.Result, l *MemberAddLogic) (*ums.MemberAddResp, error) {
	//3.1生成token
	userId, _ := result.LastInsertId()

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := l.createJwtToken(l.svcCtx.Config.JWT.AccessSecret, now, accessExpire, userId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	//3.2返回数据
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

//创建token
func (l *MemberAddLogic) createJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
