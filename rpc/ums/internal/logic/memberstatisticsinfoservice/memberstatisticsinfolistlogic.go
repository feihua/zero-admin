package memberstatisticsinfoservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberStatisticsInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberStatisticsInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberStatisticsInfoListLogic {
	return &MemberStatisticsInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberStatisticsInfoListLogic) MemberStatisticsInfoList(in *umsclient.MemberStatisticsInfoListReq) (*umsclient.MemberStatisticsInfoListResp, error) {
	q := query.UmsMemberStatisticsInfo.WithContext(l.ctx)
	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员统计列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberStatisticsInfoListData
	for _, item := range result {

		list = append(list, &umsclient.MemberStatisticsInfoListData{
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

	logc.Infof(l.ctx, "查询会员统计列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberStatisticsInfoListResp{
		Total: count,
		List:  list,
	}, nil

}
