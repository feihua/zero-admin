package memberstatisticsinfoservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberStatisticsInfoLogic 添加会员统计信息
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
*/
type AddMemberStatisticsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberStatisticsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberStatisticsInfoLogic {
	return &AddMemberStatisticsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberStatisticsInfo 添加会员统计信息
func (l *AddMemberStatisticsInfoLogic) AddMemberStatisticsInfo(in *umsclient.AddMemberStatisticsInfoReq) (*umsclient.AddMemberStatisticsInfoResp, error) {
	q := query.UmsMemberStatisticsInfo

	item := &model.UmsMemberStatisticsInfo{
		MemberID:            in.MemberId,            //
		ConsumeAmount:       in.ConsumeAmount,       // 累计消费金额
		OrderCount:          in.OrderCount,          // 订单数量
		CouponCount:         in.CouponCount,         // 优惠券数量
		CommentCount:        in.CommentCount,        // 评价数
		ReturnOrderCount:    in.ReturnOrderCount,    // 退货数量
		LoginCount:          in.LoginCount,          // 登录次数
		AttendCount:         in.AttendCount,         // 关注数量
		FansCount:           in.FansCount,           // 粉丝数量
		CollectProductCount: in.CollectProductCount, // 收藏的商品数量
		CollectSubjectCount: in.CollectSubjectCount, // 收藏的专题活动数量
		CollectTopicCount:   in.CollectTopicCount,   // 收藏的评论数量
		CollectCommentCount: in.CollectCommentCount, // 收藏的专题活动数量
		InviteFriendCount:   in.InviteFriendCount,   // 邀请好友数
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员统计信息失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加会员统计信息失败")
	}

	logc.Infof(l.ctx, "添加会员统计信息成功,参数：%+v", in)
	return &umsclient.AddMemberStatisticsInfoResp{}, nil
}
