package memberservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

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
func (l *MemberAddLogic) MemberAdd(in *umsclient.MemberAddReq) (*umsclient.MemberAddResp, error) {

	q := query.UmsMember
	count, _ := q.WithContext(l.ctx).Where(q.Username.Eq(in.Username)).Or(q.Phone.Eq(in.Phone)).Count()
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
	token, err := createJwtToken(secret, in.Username, in.Phone, accessExpire, id)

	if err != nil {
		logc.Errorf(l.ctx, "生成token失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &umsclient.MemberAddResp{
		Token: token,
	}, nil
}

func insertMember(in *umsclient.MemberAddReq, l *MemberAddLogic) (int64, error) {
	member := &model.UmsMember{
		MemberLevelID:         4, //默认是普通会员
		Username:              in.Username,
		Password:              in.Password,
		Nickname:              in.Username,
		Phone:                 in.Phone,
		Status:                0,
		CreateTime:            time.Now(),
		Icon:                  nil,
		Gender:                nil,
		Birthday:              nil,
		City:                  nil,
		Job:                   nil,
		PersonalizedSignature: nil,
		SourceType:            0,
		Integration:           0,
		Growth:                0,
		LuckeyCount:           0,
		HistoryIntegration:    0,
	}
	err := query.UmsMember.WithContext(l.ctx).Create(member)

	if err != nil {
		return 0, err
	}

	return member.ID, nil
}
