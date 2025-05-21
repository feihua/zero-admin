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

// QueryMemberInfoListLogic 查询会员信息列表
/*
Author: LiuFeiHua
Date: 2025/05/21 14:18:26
*/
type QueryMemberInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberInfoListLogic {
	return &QueryMemberInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberInfoList 查询会员信息列表
func (l *QueryMemberInfoListLogic) QueryMemberInfoList(req *types.QueryMemberInfoListReq) (resp *types.QueryMemberInfoListResp, err error) {
	result, err := l.svcCtx.MemberInfoService.QueryMemberInfoList(l.ctx, &umsclient.QueryMemberInfoListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		MemberId:  req.MemberId,  // 会员ID
		Nickname:  req.Nickname,  // 昵称
		Mobile:    req.Mobile,    // 手机号码
		Source:    req.Source,    // 注册来源：0-PC，1-APP，2-小程序
		IsEnabled: req.IsEnabled, // 是否启用：0-禁用，1-启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字会员信息列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryMemberInfoListData

	for _, detail := range result.List {
		list = append(list, &types.QueryMemberInfoListData{
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

		})
	}

	return &types.QueryMemberInfoListResp{
		Code:     "000000",
		Message:  "查询会员信息列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
