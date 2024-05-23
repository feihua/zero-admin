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

// QueryMemberStatisticsInfoListLogic 查询会员统计信息
/*
Author: LiuFeiHua
Date: 2024/5/23 9:22
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
	result, err := l.svcCtx.MemberStatisticsInfoService.MemberStatisticsInfoList(l.ctx, &umsclient.MemberStatisticsInfoListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员统计信息列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员统计信息失败")
	}

	var list []*types.ListMemberStatisticsInfoData

	for _, item := range result.List {
		list = append(list, &types.ListMemberStatisticsInfoData{
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
		})
	}

	return &types.ListMemberStatisticsInfoResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员统计信息成功",
	}, nil
}
