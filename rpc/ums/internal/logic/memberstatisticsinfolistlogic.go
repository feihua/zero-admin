package logic

import (
	"context"
	"encoding/json"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"
	"zero-admin/rpc/ums/umsclient"

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
	// todo: add your logic here and delete this line
	all, err := l.svcCtx.UmsMemberStatisticsInfoModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberStatisticsInfoModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员统计列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询会员统计列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.MemberStatisticsInfoListResp{
		Total: count,
		List:  list,
	}, nil
}
