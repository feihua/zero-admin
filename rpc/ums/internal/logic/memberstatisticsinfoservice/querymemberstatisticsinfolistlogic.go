package memberstatisticsinfoservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberStatisticsInfoListLogic 查询会员统计信息列表
type QueryMemberStatisticsInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberStatisticsInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberStatisticsInfoListLogic {
	return &QueryMemberStatisticsInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberStatisticsInfoList 查询会员统计信息列表
func (l *QueryMemberStatisticsInfoListLogic) QueryMemberStatisticsInfoList(in *umsclient.QueryMemberStatisticsInfoListReq) (*umsclient.QueryMemberStatisticsInfoListResp, error) {
	q := query.UmsMemberStatisticsInfo.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(query.UmsMemberStatisticsInfo.MemberID.Eq(in.MemberId))
	}
	offset := (in.PageNum - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员统计信息列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员统计信息列表失败")
	}

	var list []*umsclient.QueryMemberStatisticsInfoListData
	for _, item := range result {

		list = append(list, &umsclient.QueryMemberStatisticsInfoListData{
			Id:                  item.ID,                                   //
			MemberId:            item.MemberID,                             //
			ConsumeAmount:       item.ConsumeAmount,                        // 累计消费金额
			OrderCount:          item.OrderCount,                           // 订单数量
			CouponCount:         item.CouponCount,                          // 优惠券数量
			CommentCount:        item.CommentCount,                         // 评价数
			ReturnOrderCount:    item.ReturnOrderCount,                     // 退货数量
			LoginCount:          item.LoginCount,                           // 登录次数
			AttendCount:         item.AttendCount,                          // 关注数量
			FansCount:           item.FansCount,                            // 粉丝数量
			CollectProductCount: item.CollectProductCount,                  // 收藏的商品数量
			CollectSubjectCount: item.CollectSubjectCount,                  // 收藏的专题活动数量
			CollectTopicCount:   item.CollectTopicCount,                    // 收藏的评论数量
			CollectCommentCount: item.CollectCommentCount,                  // 收藏的专题活动数量
			InviteFriendCount:   item.InviteFriendCount,                    // 邀请好友数
			RecentOrderTime:     time_util.TimeToStr(item.RecentOrderTime), // 最后一次下订单时间
		})
	}

	return &umsclient.QueryMemberStatisticsInfoListResp{
		Total: count,
		List:  list,
	}, nil

}
