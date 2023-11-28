package memberservicelogic

import (
	"context"
	"encoding/json"
	"errors"
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
	//根据用户名查询账号
	member, _ := l.svcCtx.UmsMemberModel.FindMemberByNameOrPhone(l.ctx, in.Username, in.Phone)
	if member == nil {
		logx.WithContext(l.ctx).Errorf("账号不存在,参数:%s", in.Username)
		return nil, errors.New("账号不存在")
	}

	//判断密码
	if member.Password != in.Password {
		logx.WithContext(l.ctx).Errorf("账号密码不对,参数:%s", in.Password)
		return nil, errors.New("账号密码不对")
	}

	//2.添加登录日志
	insertLoginLog(l, member, in.Ip)

	//3.返回数据
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	secret := l.svcCtx.Config.JWT.AccessSecret
	token, err := createJwtToken(secret, member.Username, member.Phone, accessExpire, member.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	return &umsclient.MemberLoginResp{
		Token: token,
	}, nil
}

//插入登录日志
func insertLoginLog(l *MemberLoginLogic, m *umsmodel.UmsMember, Ip string) {
	memberLoginLog := &umsmodel.UmsMemberLoginLog{
		MemberId:   m.Id,
		CreateTime: time.Now(),
		Ip:         Ip,
		City:       m.City.String,
		LoginType:  0,
		Province:   "",
	}
	_, _ = l.svcCtx.UmsMemberLoginLogModel.Insert(l.ctx, memberLoginLog)
}
