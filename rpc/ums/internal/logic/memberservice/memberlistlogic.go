package memberservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberListLogic {
	return &MemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberListLogic) MemberList(in *umsclient.MemberListReq) (*umsclient.MemberListResp, error) {
	q := query.UmsMember.WithContext(l.ctx)
	if len(in.Username) > 0 {
		q = q.Where(query.UmsMember.MemberName.Like("%" + in.Username + "%"))
	}
	if len(in.Phone) > 0 {
		q = q.Where(query.UmsMember.Phone.Like("%" + in.Phone + "%"))
	}

	if in.Status != 2 {
		q = q.Where(query.UmsMember.MemberStatus.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}
	var list []*umsclient.MemberListData
	for _, member := range result {

		list = append(list, &umsclient.MemberListData{
			Id:                    member.ID,
			MemberLevelId:         member.MemberLevelID,
			Username:              member.MemberName,
			Nickname:              member.Nickname,
			Phone:                 member.Phone,
			Status:                member.MemberStatus,
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
			LuckeyCount:           member.LotteryCount,
			HistoryIntegration:    member.HistoryIntegration,
		})
	}

	logc.Infof(l.ctx, "查询会员列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberListResp{
		Total: count,
		List:  list,
	}, nil
}
