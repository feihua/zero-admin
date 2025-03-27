package memberservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberDetailLogic 查询会员表详情
/*
Author: LiuFeiHua
Date: 2024/6/11 13:59
*/
type QueryMemberDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberDetailLogic {
	return &QueryMemberDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberDetail 查询会员表详情
func (l *QueryMemberDetailLogic) QueryMemberDetail(in *umsclient.QueryMemberDetailReq) (*umsclient.QueryMemberDetailResp, error) {
	q := query.UmsMember
	member, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员登录失败")
	}

	memberLevel, err := query.UmsMemberLevel.WithContext(l.ctx).Where(query.UmsMemberLevel.ID.Eq(member.MemberLevelID)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员等级不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员登录失败,会员等级不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员等级异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员登录失败")
	}

	resp := &umsclient.QueryMemberDetailResp{
		Id:                    member.ID,
		MemberLevelId:         member.MemberLevelID,
		MemberName:            member.MemberName,
		Nickname:              member.Nickname,
		Phone:                 member.Phone,
		MemberStatus:          member.MemberStatus,
		CreateTime:            time_util.TimeToStr(member.CreateTime),
		Icon:                  member.Icon,
		Gender:                member.Gender,
		Birthday:              time_util.TimeToString(member.Birthday),
		City:                  member.City,
		Job:                   member.Job,
		PersonalizedSignature: member.PersonalizedSignature,
		SourceType:            member.SourceType,
		Integration:           member.Integration,
		Growth:                member.Growth,
		LotteryCount:          member.LotteryCount,
		HistoryIntegration:    member.HistoryIntegration,
		MemberLevel:           memberLevel.LevelName,
	}

	return resp, nil
}
