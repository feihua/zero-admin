package member_statistics

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberStatisticsInfoListLogic 查询会员统计信息
/*
Author: LiuFeiHua
Date: 2025/5/23 14:36
*/
type QueryMemberStatisticsInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberStatisticsInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberStatisticsInfoListLogic {
	return &QueryMemberStatisticsInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberStatisticsInfoList 查询会员统计信息
func (l *QueryMemberStatisticsInfoListLogic) QueryMemberStatisticsInfoList(req *types.ListMemberStatisticsInfoReq) (resp *types.ListMemberStatisticsInfoResp, err error) {
	result, err := l.svcCtx.MemberStatisticsInfoService.QueryMemberStatisticsInfoList(l.ctx, &umsclient.QueryMemberStatisticsInfoListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		MemberId: req.MemberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员统计信息列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员统计信息失败")
	}

	var list []*types.QueryMemberStatisticsInfoListData

	for _, item := range result.List {
		list = append(list, &types.QueryMemberStatisticsInfoListData{
			Id:                  item.Id,                  //
			MemberId:            item.MemberId,            //
			ConsumeAmount:       item.ConsumeAmount,       // 累计消费金额
			OrderCount:          item.OrderCount,          // 订单数量
			CouponCount:         item.CouponCount,         // 优惠券数量
			CommentCount:        item.CommentCount,        // 评价数
			ReturnOrderCount:    item.ReturnOrderCount,    // 退货数量
			LoginCount:          item.LoginCount,          // 登录次数
			AttendCount:         item.AttendCount,         // 关注数量
			FansCount:           item.FansCount,           // 粉丝数量
			CollectProductCount: item.CollectProductCount, // 收藏的商品数量
			CollectSubjectCount: item.CollectSubjectCount, // 收藏的专题活动数量
			CollectTopicCount:   item.CollectTopicCount,   // 收藏的评论数量
			CollectCommentCount: item.CollectCommentCount, // 收藏的专题活动数量
			InviteFriendCount:   item.InviteFriendCount,   // 邀请好友数
			RecentOrderTime:     item.RecentOrderTime,     // 最后一次下订单时间
		})
	}

	return &types.ListMemberStatisticsInfoResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员统计信息列表成功",
	}, nil
}
