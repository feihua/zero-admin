package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

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

func (l *MemberListLogic) MemberList(in *ums.MemberListReq) (*ums.MemberListResp, error) {
	all, err := l.svcCtx.UmsMemberModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询会员列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &ums.MemberListResp{
		Total: count,
		List:  list,
	}, nil
}
