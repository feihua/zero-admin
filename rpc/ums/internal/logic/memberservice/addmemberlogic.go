package memberservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberLogic 添加会员表
/*
Author: LiuFeiHua
Date: 2024/6/11 13:53
*/
type AddMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberLogic {
	return &AddMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMember 添加会员表
func (l *AddMemberLogic) AddMember(in *umsclient.AddMemberReq) (*umsclient.AddMemberResp, error) {
	q := query.UmsMember
	count, _ := q.WithContext(l.ctx).Where(q.MemberName.Eq(in.MemberName)).Or(q.Phone.Eq(in.Phone)).Count()
	if count > 0 {
		logc.Errorf(l.ctx, "用户名或手机号已注册,参数: %+v", in)
		return nil, errors.New("用户名或手机号已注册")
	}

	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	secret := l.svcCtx.Config.JWT.AccessSecret
	id, err := insertMember(in, l)
	if err != nil {
		logc.Errorf(l.ctx, "新增会员失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("新增会员失败")
	}
	token, err := createJwtToken(secret, in.MemberName, in.Phone, accessExpire, id)

	if err != nil {
		logc.Errorf(l.ctx, "生成token失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &umsclient.AddMemberResp{
		Token: token,
	}, nil
}

func insertMember(in *umsclient.AddMemberReq, l *AddMemberLogic) (int64, error) {
	member := &model.UmsMember{
		MemberLevelID:      4, // 默认是普通会员
		MemberName:         in.MemberName,
		Password:           in.Password,
		Nickname:           in.MemberName,
		Phone:              in.Phone,
		MemberStatus:       0,
		CreateTime:         time.Now(),
		SourceType:         0,
		Integration:        0,
		Growth:             0,
		LotteryCount:       0,
		HistoryIntegration: 0,
	}
	err := query.UmsMember.WithContext(l.ctx).Create(member)

	if err != nil {
		return 0, err
	}

	return member.ID, nil
}
