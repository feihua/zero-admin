package memberstatisticsinfoservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberStatisticsInfoLogic 添加会员统计信息
/*
Author: LiuFeiHua
Date: 2024/6/11 13:48
*/
type AddMemberStatisticsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberStatisticsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberStatisticsInfoLogic {
	return &AddMemberStatisticsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberStatisticsInfo 添加会员统计信息
func (l *AddMemberStatisticsInfoLogic) AddMemberStatisticsInfo(in *umsclient.AddMemberStatisticsInfoReq) (*umsclient.AddMemberStatisticsInfoResp, error) {
	err := query.UmsMemberStatisticsInfo.WithContext(l.ctx).Create(&model.UmsMemberStatisticsInfo{
		MemberID:            in.MemberId,
		ConsumeAmount:       in.ConsumeAmount,
		OrderCount:          in.OrderCount,
		CouponCount:         in.CouponCount,
		CommentCount:        in.CommentCount,
		ReturnOrderCount:    in.ReturnOrderCount,
		LoginCount:          in.LoginCount,
		AttendCount:         in.AttendCount,
		FansCount:           in.FansCount,
		CollectProductCount: in.CollectProductCount,
		CollectSubjectCount: in.CollectSubjectCount,
		CollectTopicCount:   in.CollectTopicCount,
		CollectCommentCount: in.CollectCommentCount,
		InviteFriendCount:   in.InviteFriendCount,
		RecentOrderTime:     time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddMemberStatisticsInfoResp{}, nil
}
