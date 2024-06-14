package memberstatisticsinfoservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

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

	offset := (in.PageNum - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员统计信息列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.QueryMemberStatisticsInfoListData
	for _, item := range result {

		list = append(list, &umsclient.QueryMemberStatisticsInfoListData{
			Id:                  item.ID,
			MemberId:            item.MemberID,
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
			RecentOrderTime:     item.RecentOrderTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &umsclient.QueryMemberStatisticsInfoListResp{
		Total: count,
		List:  list,
	}, nil

}
