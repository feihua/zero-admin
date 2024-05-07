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

// MemberStatisticsInfoUpdateLogic 更新会员统计信息
/*
Author: LiuFeiHua
Date: 2024/5/7 10:43
*/
type MemberStatisticsInfoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberStatisticsInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberStatisticsInfoUpdateLogic {
	return &MemberStatisticsInfoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberStatisticsInfoUpdate 更新会员统计信息
func (l *MemberStatisticsInfoUpdateLogic) MemberStatisticsInfoUpdate(in *umsclient.MemberStatisticsInfoUpdateReq) (*umsclient.MemberStatisticsInfoUpdateResp, error) {
	_, err := query.UmsMemberStatisticsInfo.WithContext(l.ctx).Updates(&model.UmsMemberStatisticsInfo{
		ID:                  in.Id,
		MemberID:            in.MemberId,
		ConsumeAmount:       float64(in.ConsumeAmount),
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

	return &umsclient.MemberStatisticsInfoUpdateResp{}, nil
}
