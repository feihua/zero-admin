package memberservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

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

	if err != nil {
		logc.Errorf(l.ctx, "根据会员id查询会员信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	resp := &umsclient.QueryMemberDetailResp{
		Id:                    member.ID,
		MemberLevelId:         member.MemberLevelID,
		MemberName:            member.MemberName,
		Nickname:              member.Nickname,
		Phone:                 member.Phone,
		MemberStatus:          member.MemberStatus,
		CreateTime:            member.CreateTime.Format("2006-01-02 15:04:05"),
		Icon:                  member.Icon,
		Gender:                member.Gender,
		Birthday:              member.Birthday.Format("2006-01-02 15:04:05"),
		City:                  member.City,
		Job:                   member.Job,
		PersonalizedSignature: member.PersonalizedSignature,
		SourceType:            member.SourceType,
		Integration:           member.Integration,
		Growth:                member.Growth,
		LotteryCount:          member.LotteryCount,
		HistoryIntegration:    member.HistoryIntegration,
	}

	return resp, nil
}
