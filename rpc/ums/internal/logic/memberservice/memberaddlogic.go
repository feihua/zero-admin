package memberservicelogic

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/feihua/zero-admin/rpc/model/umsmodel"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
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

	username := in.Username
	member, _ := l.svcCtx.UmsMemberModel.FindOneByUsername(l.ctx, username)
	if member != nil {
		logx.WithContext(l.ctx).Errorf("用户名已注册,参数Username:%s", username)
		return nil, errors.New("用户名已注册")
	}

	p := in.Phone
	phone, _ := l.svcCtx.UmsMemberModel.FindOneByPhone(l.ctx, p)
	if phone != nil {
		logx.WithContext(l.ctx).Errorf("手机号已注册,参数Phone:%s", p)
		return nil, errors.New("手机号已注册")
	}

	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	secret := l.svcCtx.Config.JWT.AccessSecret
	token, err := createJwtToken(secret, username, p, accessExpire, insertMember(in, l))

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	return &umsclient.MemberAddResp{
		Token: token,
	}, nil
}

func insertMember(in *umsclient.MemberAddReq, l *MemberAddLogic) int64 {
	result, _ := l.svcCtx.UmsMemberModel.Insert(l.ctx, &umsmodel.UmsMember{
		MemberLevelId:         4, //默认是普通会员
		Username:              in.Username,
		Password:              in.Password,
		Nickname:              in.Username,
		Phone:                 in.Phone,
		Status:                0,
		CreateTime:            time.Now(),
		Icon:                  sql.NullString{},
		Gender:                sql.NullInt64{},
		Birthday:              sql.NullTime{},
		City:                  sql.NullString{},
		Job:                   sql.NullString{},
		PersonalizedSignature: sql.NullString{},
		SourceType:            0,
		Integration:           0,
		Growth:                0,
		LuckeyCount:           0,
		HistoryIntegration:    0,
	})
	memberId, _ := result.LastInsertId()
	return memberId
}
