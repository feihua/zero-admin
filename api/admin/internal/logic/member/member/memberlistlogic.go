package member

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

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
	resp, err := l.svcCtx.MemberService.MemberList(l.ctx, &umsclient.MemberListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Username: req.Username,
		Phone:    req.Phone,
		Status:   req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员信息列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员信息失败")
	}

	var list []*types.ListMemberData

	for _, member := range resp.List {
		list = append(list, &types.ListMemberData{
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
