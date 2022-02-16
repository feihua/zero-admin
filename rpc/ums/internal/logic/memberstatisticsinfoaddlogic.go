package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberStatisticsInfoAddLogic) MemberStatisticsInfoAdd(in *ums.MemberStatisticsInfoAddReq) (*ums.MemberStatisticsInfoAddResp, error) {
	RecentOrderTime, _ := time.Parse("2006-01-02 15:04:05", in.RecentOrderTime)
	_, err := l.svcCtx.UmsMemberStatisticsInfoModel.Insert(umsmodel.UmsMemberStatisticsInfo{
		MemberId:            in.MemberId,
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
		RecentOrderTime:     RecentOrderTime,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberStatisticsInfoAddResp{}, nil
}
