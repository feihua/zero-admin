package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberStatisticsInfoAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberStatisticsInfoAddLogic {
	return MemberStatisticsInfoAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoAddLogic) MemberStatisticsInfoAdd(req types.AddMemberStatisticsInfoReq) (*types.AddMemberStatisticsInfoResp, error) {
	_, err := l.svcCtx.Ums.MemberStatisticsInfoAdd(l.ctx, &umsclient.MemberStatisticsInfoAddReq{
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员统计信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员统计信息失败")
	}

	return &types.AddMemberStatisticsInfoResp{
		Code:    "000000",
		Message: "添加会员统计信息成功",
	}, nil
}
