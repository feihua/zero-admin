package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	_, err := l.svcCtx.MemberStatisticsInfoService.MemberStatisticsInfoUpdate(l.ctx, &umsclient.MemberStatisticsInfoUpdateReq{
		Id:                  req.Id,
		MemberId:            req.MemberId,
		ConsumeAmount:       req.ConsumeAmount,
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
		logc.Errorf(l.ctx, "更新会员统计信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员统计信息失败")
	}

	return &types.UpdateMemberStatisticsInfoResp{
		Code:    "000000",
		Message: "更新会员统计信息成功",
	}, nil
}
