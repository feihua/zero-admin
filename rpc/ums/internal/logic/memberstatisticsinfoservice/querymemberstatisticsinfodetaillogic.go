package memberstatisticsinfoservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberStatisticsInfoDetailLogic 查询会员统计信息详情
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
*/
type QueryMemberStatisticsInfoDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberStatisticsInfoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberStatisticsInfoDetailLogic {
	return &QueryMemberStatisticsInfoDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberStatisticsInfoDetail 查询会员统计信息详情
func (l *QueryMemberStatisticsInfoDetailLogic) QueryMemberStatisticsInfoDetail(in *umsclient.QueryMemberStatisticsInfoDetailReq) (*umsclient.QueryMemberStatisticsInfoDetailResp, error) {
	item, err := query.UmsMemberStatisticsInfo.WithContext(l.ctx).Where(query.UmsMemberStatisticsInfo.ID.Eq(in.MemberId)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员统计信息详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员统计信息详情失败")
	}

	data := &umsclient.QueryMemberStatisticsInfoDetailResp{
		Id:                  item.ID,                                            //
		MemberId:            item.MemberID,                                      //
		ConsumeAmount:       item.ConsumeAmount,                                 // 累计消费金额
		OrderCount:          item.OrderCount,                                    // 订单数量
		CouponCount:         item.CouponCount,                                   // 优惠券数量
		CommentCount:        item.CommentCount,                                  // 评价数
		ReturnOrderCount:    item.ReturnOrderCount,                              // 退货数量
		LoginCount:          item.LoginCount,                                    // 登录次数
		AttendCount:         item.AttendCount,                                   // 关注数量
		FansCount:           item.FansCount,                                     // 粉丝数量
		CollectProductCount: item.CollectProductCount,                           // 收藏的商品数量
		CollectSubjectCount: item.CollectSubjectCount,                           // 收藏的专题活动数量
		CollectTopicCount:   item.CollectTopicCount,                             // 收藏的评论数量
		CollectCommentCount: item.CollectCommentCount,                           // 收藏的专题活动数量
		InviteFriendCount:   item.InviteFriendCount,                             // 邀请好友数
		RecentOrderTime:     item.RecentOrderTime.Format("2006-01-02 15:04:05"), // 最后一次下订单时间
	}

	return data, nil
}
