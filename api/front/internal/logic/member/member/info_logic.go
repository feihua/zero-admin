package member

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// InfoLogic 获取会员个人信息
/*
Author: LiuFeiHua
Date: 2025/6/19 11:08
*/
type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Info 获取会员个人信息
func (l *InfoLogic) Info() (resp *types.InfoResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	detail, err := l.svcCtx.MemberService.QueryMemberInfoDetail(l.ctx, &umsclient.QueryMemberInfoDetailReq{MemberId: memberId})

	if err != nil {
		logc.Errorf(l.ctx, "获取个人信息失败,参数memberId：%d,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.InfoResp{
		Code:    0,
		Message: "查询会员信息",
		Data: types.MemberData{
			Id:           detail.Id,                   // 主键ID
			MemberId:     detail.MemberId,             // 会员ID
			LevelId:      detail.LevelId,              // 等级ID
			Nickname:     detail.Nickname,             // 昵称
			Mobile:       detail.Mobile,               // 手机号码
			Source:       detail.Source,               // 注册来源：0-PC，1-APP，2-小程序
			Avatar:       detail.Avatar,               // 头像
			Signature:    detail.Signature,            // 个性签名
			Gender:       detail.Gender,               // 性别：0-未知，1-男，2-女
			Birthday:     detail.Birthday,             // 生日
			GrowthPoint:  detail.GrowthPoint,          // 成长值
			Points:       detail.Points,               // 积分
			TotalPoints:  detail.TotalPoints,          // 累计获得积分
			SpendAmount:  float64(detail.SpendAmount), // 累计消费金额
			OrderCount:   detail.OrderCount,           // 订单数
			CouponCount:  detail.CouponCount,          // 优惠券数量
			CommentCount: detail.CommentCount,         // 评价数
			ReturnCount:  detail.ReturnCount,          // 退货数
			LotteryTimes: detail.LotteryTimes,         // 剩余抽奖次数
			LastLogin:    detail.LastLogin,            // 最后登录
		},
	}, nil
}
