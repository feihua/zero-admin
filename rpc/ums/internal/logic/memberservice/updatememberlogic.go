package memberservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberLogic 更新会员表
/*
Author: LiuFeiHua
Date: 2024/6/11 13:58
*/
type UpdateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLogic {
	return &UpdateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMember 更新会员表
func (l *UpdateMemberLogic) UpdateMember(in *umsclient.UpdateMemberReq) (*umsclient.UpdateMemberResp, error) {
	birthday, _ := time.Parse("2006-01-02 15:04:05", in.Birthday)
	_, err := query.UmsMember.WithContext(l.ctx).Updates(&model.UmsMember{
		ID:                    in.Id,
		MemberLevelID:         in.MemberLevelId,
		MemberName:            in.MemberName,
		Nickname:              in.Nickname,
		Phone:                 in.Phone,
		MemberStatus:          in.MemberStatus,
		Icon:                  &in.Icon,
		Gender:                &in.Gender,
		Birthday:              &birthday,
		City:                  &in.City,
		Job:                   &in.Job,
		PersonalizedSignature: &in.PersonalizedSignature,
		SourceType:            in.SourceType,
		Integration:           in.Integration,
		Growth:                in.Growth,
		LotteryCount:          in.LotteryCount,
		HistoryIntegration:    in.HistoryIntegration,
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberResp{}, nil
}
