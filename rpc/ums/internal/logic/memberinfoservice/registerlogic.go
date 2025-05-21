package memberinfoservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// RegisterLogic 注册会员信息
/*
Author: LiuFeiHua
Date: 2025/5/21 15:13
*/
type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register 注册会员信息
func (l *RegisterLogic) Register(in *umsclient.RegisterReq) (*umsclient.RegisterResp, error) {
	q := query.UmsMemberInfo
	count, _ := q.WithContext(l.ctx).Where(q.Mobile.Eq(in.Mobile)).Count()
	if count > 0 {
		logc.Errorf(l.ctx, "手机号已注册,参数: %+v", in)
		return nil, errors.New("手机号已注册")
	}

	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	secret := l.svcCtx.Config.JWT.AccessSecret
	id, err := insertMember(in, l)
	if err != nil {
		logc.Errorf(l.ctx, "新增会员失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("新增会员失败")
	}
	token, err := createJwtToken(secret, in.Nickname, in.Mobile, accessExpire, id)

	if err != nil {
		logc.Errorf(l.ctx, "生成token失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &umsclient.RegisterResp{
		Token: token,
	}, nil
}

func insertMember(in *umsclient.RegisterReq, l *RegisterLogic) (int64, error) {
	last, err := query.UmsMemberInfo.WithContext(l.ctx).Last()
	member := &model.UmsMemberInfo{
		MemberID: last.MemberID + 1, // 会员ID
		LevelID:  1,                 // 等级ID
		Nickname: in.Nickname,       // 昵称
		Mobile:   in.Mobile,         // 手机号码
		Source:   in.Source,         // 注册来源：0-PC，1-APP，2-小程序
		Password: in.Password,       // 密码
	}
	err = query.UmsMemberInfo.WithContext(l.ctx).Create(member)

	if err != nil {
		return 0, err
	}

	return member.MemberID, nil
}
