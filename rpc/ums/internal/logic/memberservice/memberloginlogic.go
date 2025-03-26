package memberservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
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

// MemberLogin
// 1.校验参数
// 2.添加登录日志
// 3.返回数据
func (l *MemberLoginLogic) MemberLogin(in *umsclient.MemberLoginReq) (*umsclient.MemberLoginResp, error) {

	q := query.UmsMember
	// 1.校验参数
	member, err := q.WithContext(l.ctx).Where(q.MemberName.Eq(in.Account)).Or(q.Phone.Eq(in.Account)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "账号不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("账号不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员异常")
	}

	// 判断密码
	if member.Password != in.Password {
		logc.Errorf(l.ctx, "账号密码不对,请求参数：%+v", in)
		return nil, errors.New("账号密码不对")
	}

	// 2.添加登录日志
	log := &model.UmsMemberLoginLog{
		MemberID:   member.ID,
		CreateTime: time.Now(),
		MemberIP:   in.Ip,
		City:       "todo",
		LoginType:  in.LoginType,
		Province:   "todo",
	}
	err = query.UmsMemberLoginLog.WithContext(l.ctx).Create(log)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员登录日志失败,参数：%+v,异常:%s", log, err.Error())
		// 为了兼容，这里不返回错误
	}

	// 3.返回数据
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	secret := l.svcCtx.Config.JWT.AccessSecret
	token, err := createJwtToken(secret, member.MemberName, member.Phone, accessExpire, member.ID)

	if err != nil {
		logc.Errorf(l.ctx, "生成token失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("生成token失败")
	}

	return &umsclient.MemberLoginResp{
		Token: token,
	}, nil
}
