package memberstatisticsinfoservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberStatisticsInfoDetailLogic 查询会员统计信息详情
/*
Author: LiuFeiHua
Date: 2024/6/11 13:50
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
	q := query.UmsMemberStatisticsInfo
	item, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员统计列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	data := &umsclient.QueryMemberStatisticsInfoDetailResp{
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
	}

	return data, nil

}
