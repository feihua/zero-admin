package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *MemberStatisticsInfoListLogic) MemberStatisticsInfoList(in *ums.MemberStatisticsInfoListReq) (*ums.MemberStatisticsInfoListResp, error) {
	all, _ := l.svcCtx.UmsMemberStatisticsInfoModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberStatisticsInfoModel.Count()

	var list []*ums.MemberStatisticsInfoListData
	for _, item := range *all {

		list = append(list, &ums.MemberStatisticsInfoListData{
			Id:                  item.Id,
			MemberId:            item.MemberId,
			ConsumeAmount:       int64(item.ConsumeAmount),
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

	fmt.Println(list)
	return &ums.MemberStatisticsInfoListResp{
		Total: count,
		List:  list,
	}, nil

}
