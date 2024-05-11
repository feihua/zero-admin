package memberstatisticsinfoservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberStatisticsInfoAddLogic 会员统计信息
/*
Author: LiuFeiHua
Date: 2024/5/7 10:43
*/
type MemberStatisticsInfoAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberStatisticsInfoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberStatisticsInfoAddLogic {
	return &MemberStatisticsInfoAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberStatisticsInfoAdd 添加会员统计信息
func (l *MemberStatisticsInfoAddLogic) MemberStatisticsInfoAdd(in *umsclient.MemberStatisticsInfoAddReq) (*umsclient.MemberStatisticsInfoAddResp, error) {
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

	return &umsclient.MemberStatisticsInfoAddResp{}, nil
}
