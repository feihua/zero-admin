package logic

import (
	"context"
	"time"

	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberStatisticsInfoUpdateLogic) MemberStatisticsInfoUpdate(in *umsclient.MemberStatisticsInfoUpdateReq) (*umsclient.MemberStatisticsInfoUpdateResp, error) {
	// todo: add your logic here and delete this line
	RecentOrderTime, _ := time.Parse("2006-01-02 15:04:05", in.RecentOrderTime)
	err := l.svcCtx.UmsMemberStatisticsInfoModel.Update(l.ctx, &umsmodel.UmsMemberStatisticsInfo{
		Id:                  in.Id,
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
	return &umsclient.MemberStatisticsInfoUpdateResp{}, nil
}
