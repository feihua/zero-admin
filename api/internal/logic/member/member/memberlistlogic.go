package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberListLogic {
	return MemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberListLogic) MemberList(req types.ListMemberReq) (*types.ListMemberResp, error) {
	resp, err := l.svcCtx.Ums.MemberList(l.ctx, &umsclient.MemberListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询会员信息列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询会员信息失败")
	}

	var list []*types.ListtMemberData

	for _, member := range resp.List {
		list = append(list, &types.ListtMemberData{
			Id:                    member.Id,
			MemberLevelId:         member.MemberLevelId,
			Username:              member.Username,
			Password:              member.Password,
			Nickname:              member.Nickname,
			Phone:                 member.Phone,
			Status:                member.Status,
			CreateTime:            member.CreateTime,
			Icon:                  member.Icon,
			Gender:                member.Gender,
			Birthday:              member.Birthday,
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

	return &types.ListMemberResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询会员信息成功",
	}, nil
}
