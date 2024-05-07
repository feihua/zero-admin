package memberservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberByIdLogic
/*
Author: LiuFeiHua
Date: 2023/11/28 15:07
*/
type QueryMemberByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberByIdLogic {
	return &QueryMemberByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberById 获取个人信息
func (l *QueryMemberByIdLogic) QueryMemberById(in *umsclient.MemberByIdReq) (*umsclient.MemberListData, error) {
	q := query.UmsMember
	member, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据会员id查询会员信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	resp := &umsclient.MemberListData{
		Id:                    member.ID,
		MemberLevelId:         member.MemberLevelID,
		Username:              member.Username,
		Password:              member.Password,
		Nickname:              member.Nickname,
		Phone:                 member.Phone,
		Status:                member.Status,
		CreateTime:            member.CreateTime.Format("2006-01-02 15:04:05"),
		Icon:                  *member.Icon,
		Gender:                *member.Gender,
		Birthday:              member.Birthday.Format("2006-01-02 15:04:05"),
		City:                  *member.City,
		Job:                   *member.Job,
		PersonalizedSignature: *member.PersonalizedSignature,
		SourceType:            member.SourceType,
		Integration:           member.Integration,
		Growth:                member.Growth,
		LuckeyCount:           member.LuckeyCount,
		HistoryIntegration:    member.HistoryIntegration,
	}

	logc.Infof(l.ctx, "根据会员id查询会员信息,参数：%+v,响应：%+v", in, resp)

	return resp, nil
}
