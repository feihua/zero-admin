package statistics

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberStatisticsInfoDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberStatisticsInfoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberStatisticsInfoDetailLogic {
	return &QueryMemberStatisticsInfoDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMemberStatisticsInfoDetailLogic) QueryMemberStatisticsInfoDetail(req *types.QueryMemberStatisticsInfoDetailReq) (resp *types.QueryMemberStatisticsInfoDetailResp, err error) {
	item, err := l.svcCtx.MemberStatisticsInfoService.QueryMemberStatisticsInfoDetail(l.ctx, &umsclient.QueryMemberStatisticsInfoDetailReq{
		MemberId: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员统计信息列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员统计信息失败")
	}

	data := types.QueryMemberStatisticsInfoDetailData{
		Id:                  item.Id,
		MemberId:            item.MemberId,
		ConsumeAmount:       item.ConsumeAmount,
		OrderCount:          item.OrderCount,
		CouponCount:         item.CouponCount,
		CommentCount:        item.CommentCount,
		ReturnOrderCount:    item.ReturnOrderCount,
		LoginCount:          item.LoginCount,
		AttendCount:         item.AttendCount,
		FansCount:           item.FansCount,
		CollectProductCount: item.CollectProductCount,
		CollectSubjectCount: item.CollectSubjectCount,
		CollectTopicCount:   item.CollectTopicCount,
		CollectCommentCount: item.CollectCommentCount,
		InviteFriendCount:   item.InviteFriendCount,
		RecentOrderTime:     item.RecentOrderTime,
	}

	return &types.QueryMemberStatisticsInfoDetailResp{
		Code:    "000000",
		Message: "查询会员统计信息成功",
		Data:    data,
	}, nil
}
