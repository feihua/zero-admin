package memberservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberListLogic 查询会员表列表
/*
Author: LiuFeiHua
Date: 2024/6/11 13:56
*/
type QueryMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberListLogic {
	return &QueryMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberList 查询会员表列表
func (l *QueryMemberListLogic) QueryMemberList(in *umsclient.QueryMemberListReq) (*umsclient.QueryMemberListResp, error) {
	q := query.UmsMember.WithContext(l.ctx)
	if len(in.MemberName) > 0 {
		q = q.Where(query.UmsMember.MemberName.Like("%" + in.MemberName + "%"))
	}
	if len(in.Phone) > 0 {
		q = q.Where(query.UmsMember.Phone.Like("%" + in.Phone + "%"))
	}

	if in.MemberStatus != 2 {
		q = q.Where(query.UmsMember.MemberStatus.Eq(in.MemberStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}
	var list []*umsclient.MemberListData
	for _, member := range result {

		list = append(list, &umsclient.MemberListData{
			Id:                    member.ID,
			MemberLevelId:         member.MemberLevelID,
			MemberName:            member.MemberName,
			Nickname:              member.Nickname,
			Phone:                 member.Phone,
			MemberStatus:          member.MemberStatus,
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
			LotteryCount:          member.LotteryCount,
			HistoryIntegration:    member.HistoryIntegration,
		})
	}

	return &umsclient.QueryMemberListResp{
		Total: count,
		List:  list,
	}, nil
}
