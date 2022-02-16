package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberStatisticsInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberStatisticsInfoListLogic {
	return MemberStatisticsInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoListLogic) MemberStatisticsInfoList(req types.ListMemberStatisticsInfoReq) (*types.ListMemberStatisticsInfoResp, error) {
	resp, err := l.svcCtx.Ums.MemberStatisticsInfoList(l.ctx, &umsclient.MemberStatisticsInfoListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询会员统计信息列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询会员统计信息失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtMemberStatisticsInfoData

	for _, item := range resp.List {
		list = append(list, &types.ListtMemberStatisticsInfoData{
			Id:                  item.Id,
			MemberId:            item.MemberId,
			ConsumeAmount:       float64(item.ConsumeAmount),
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
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询会员统计信息成功",
	}, nil
}
