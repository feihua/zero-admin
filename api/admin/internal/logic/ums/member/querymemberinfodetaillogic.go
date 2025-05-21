package member

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberInfoDetailLogic 查询会员信息详情
/*
Author: LiuFeiHua
Date: 2025/05/21 14:18:26
*/
type QueryMemberInfoDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberInfoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberInfoDetailLogic {
	return &QueryMemberInfoDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberInfoDetail 查询会员信息详情
func (l *QueryMemberInfoDetailLogic) QueryMemberInfoDetail(req *types.QueryMemberInfoDetailReq) (resp *types.QueryMemberInfoDetailResp, err error) {

	detail, err := l.svcCtx.MemberInfoService.QueryMemberInfoDetail(l.ctx, &umsclient.QueryMemberInfoDetailReq{
		MemberId: req.MemberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员信息详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberInfoDetailData{
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
		IsEnabled:    detail.IsEnabled,            // 是否启用：0-禁用，1-启用
		CreateTime:   detail.CreateTime,           // 创建时间
		UpdateTime:   detail.UpdateTime,           // 更新时间

	}
	return &types.QueryMemberInfoDetailResp{
		Code:    "000000",
		Message: "查询会员信息成功",
		Data:    data,
	}, nil
}
