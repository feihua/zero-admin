package memberservicelogic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLoginLogic 会员登录
/*
Author: LiuFeiHua
Date: 2023/11/28 14:33
*/
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

	q := query.UmsMember
	//1.校验参数
	//根据用户名查询账号
	member, _ := q.WithContext(l.ctx).Where(q.Username.Eq(in.Account)).Or(q.Phone.Eq(in.Account)).First()
	if member == nil {
		logc.Errorf(l.ctx, "账号不存在,参数:%s", in.Account)
		return nil, errors.New("账号不存在")
	}

	//判断密码
	if member.Password != in.Password {
		logc.Errorf(l.ctx, "账号密码不对,参数:%s", in.Password)
		return nil, errors.New("账号密码不对")
	}

	//2.添加登录日志
	log := &model.UmsMemberLoginLog{
		MemberID:   member.ID,
		CreateTime: time.Now(),
		IP:         in.Ip,
		City:       *member.City,
		LoginType:  0,
		Province:   "",
	}
	_ = query.UmsMemberLoginLog.WithContext(l.ctx).Create(log)

	//3.返回数据
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	secret := l.svcCtx.Config.JWT.AccessSecret
	token, err := createJwtToken(secret, member.Username, member.Phone, accessExpire, member.ID)

	if err != nil {
		in, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "生成token失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &umsclient.MemberLoginResp{
		Token: token,
	}, nil
}
