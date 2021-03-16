package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberStatisticsInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberStatisticsInfoUpdateLogic {
	return MemberStatisticsInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoUpdateLogic) MemberStatisticsInfoUpdate(req types.UpdateMemberStatisticsInfoReq) (*types.UpdateMemberStatisticsInfoResp, error) {
	_, err := l.svcCtx.Ums.MemberStatisticsInfoUpdate(l.ctx, &umsclient.MemberStatisticsInfoUpdateReq{
		Id:                  req.Id,
		MemberId:            req.MemberId,
		ConsumeAmount:       int64(req.ConsumeAmount),
		OrderCount:          req.OrderCount,
		CouponCount:         req.CouponCount,
		CommentCount:        req.CommentCount,
		ReturnOrderCount:    req.ReturnOrderCount,
		LoginCount:          req.LoginCount,
		AttendCount:         req.AttendCount,
		FansCount:           req.FansCount,
		CollectProductCount: req.CollectProductCount,
		CollectSubjectCount: req.CollectSubjectCount,
		CollectTopicCount:   req.CollectTopicCount,
		CollectCommentCount: req.CollectCommentCount,
		InviteFriendCount:   req.InviteFriendCount,
		RecentOrderTime:     req.RecentOrderTime,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateMemberStatisticsInfoResp{
		Code:    "000000",
		Message: "",
	}, nil
}
