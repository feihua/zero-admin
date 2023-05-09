package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogic {
	return &MemberLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogic) MemberLogin(in *umsclient.MemberLoginReq) (*umsclient.MemberLoginResp, error) {

	//1.校验参数
	member, err2 := checkLoginParams(in, l)
	if err2 != nil {
		return nil, err2
	}

	//2.添加登录日志
	insertLoginLog(l, member)

	//3.返回数据
	return buildLoginResp(in, l, member)
}

//返回数据
func buildLoginResp(in *umsclient.MemberLoginReq, l *MemberLoginLogic, member *umsmodel.UmsMember) (*umsclient.MemberLoginResp, error) {
	//生成token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.JWT.AccessSecret, now, accessExpire, member.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	resp := &umsclient.MemberLoginResp{
		Nickname: in.Username,
		Token:    jwtToken,
		Icon:     "",
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(resp)
	logx.WithContext(l.ctx).Infof("登录成功,参数:%s,响应:%s", reqStr, listStr)

	return resp, nil
}

func checkLoginParams(in *umsclient.MemberLoginReq, l *MemberLoginLogic) (*umsmodel.UmsMember, error) {
	//根据用户名查询账号
	member, _ := l.svcCtx.UmsMemberModel.FindOneByUsername(l.ctx, in.Username)
	if member == nil {
		logx.WithContext(l.ctx).Errorf("账号不存在,参数:%s", in.Username)
		return nil, errors.New("账号不存在")
	}

	//判断密码
	if member.Password != in.Password {
		logx.WithContext(l.ctx).Errorf("账号密码不对,参数:%s", in.Password)
		return nil, errors.New("账号密码不对")
	}
	return member, nil
}

//插入登录日志
func insertLoginLog(l *MemberLoginLogic, m *umsmodel.UmsMember) {
	memberLoginLog := &umsmodel.UmsMemberLoginLog{
		MemberId:   m.Id,
		CreateTime: time.Now(),
		Ip:         "",
		City:       m.City.String,
		LoginType:  0,
		Province:   "",
	}
	_, _ = l.svcCtx.UmsMemberLoginLogModel.Insert(l.ctx, memberLoginLog)
}

//生成token
func (l *MemberLoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
