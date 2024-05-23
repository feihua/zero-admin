package memberservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberUpdateLogic 会员信息
type MemberUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberUpdateLogic {
	return &MemberUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberUpdate 更新会员信息
func (l *MemberUpdateLogic) MemberUpdate(in *umsclient.MemberUpdateReq) (*umsclient.MemberUpdateResp, error) {
	createTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	birthday, _ := time.Parse("2006-01-02 15:04:05", in.Birthday)
	_, err := query.UmsMember.WithContext(l.ctx).Updates(&model.UmsMember{
		ID:                    in.Id,
		MemberLevelID:         in.MemberLevelId,
		MemberName:            in.Username,
		Nickname:              in.Nickname,
		Phone:                 in.Phone,
		MemberStatus:          in.Status,
		CreateTime:            createTime,
		Icon:                  &in.Icon,
		Gender:                &in.Gender,
		Birthday:              &birthday,
		City:                  &in.City,
		Job:                   &in.Job,
		PersonalizedSignature: &in.PersonalizedSignature,
		SourceType:            in.SourceType,
		Integration:           in.Integration,
		Growth:                in.Growth,
		LotteryCount:          in.LuckeyCount,
		HistoryIntegration:    in.HistoryIntegration,
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberUpdateResp{}, nil
}
