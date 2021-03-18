package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *MemberListLogic) MemberList(in *ums.MemberListReq) (*ums.MemberListResp, error) {
	all, _ := l.svcCtx.UmsMemberModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberModel.Count()

	var list []*ums.MemberListData
	for _, member := range *all {

		list = append(list, &ums.MemberListData{
			Id:                    member.Id,
			MemberLevelId:         member.MemberLevelId,
			Username:              member.Username,
			Password:              member.Password,
			Nickname:              member.Nickname,
			Phone:                 member.Phone,
			Status:                member.Status,
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
			LuckeyCount:           member.LuckeyCount,
			HistoryIntegration:    member.HistoryIntegration,
		})
	}

	fmt.Println(list)
	return &ums.MemberListResp{
		Total: count,
		List:  list,
	}, nil
}
