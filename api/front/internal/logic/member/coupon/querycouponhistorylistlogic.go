package coupon

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponHistoryListLogic 获取会员优惠券历史列表
/*
Author: LiuFeiHua
Date: 2024/5/16 16:20
*/
type QueryCouponHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponHistoryListLogic {
	return &QueryCouponHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponHistoryList 获取会员优惠券历史列表
func (l *QueryCouponHistoryListLogic) QueryCouponHistoryList(req *types.ListCouponHistoryReq) (resp *types.ListCouponHistoryResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	historyList, err := l.svcCtx.CouponRecordService.QueryCouponRecordList(l.ctx, &smsclient.QueryCouponRecordListReq{
		PageNum:  1,
		PageSize: 100,
		MemberId: memberId,
		Status:   req.UseStatus,
	})
	if err != nil {
		logc.Errorf(l.ctx, "获取会员优惠券历史列表失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.ListCouponHistoryData

	for _, detail := range historyList.List {
		list = append(list, &types.ListCouponHistoryData{
			Id:             detail.Id,             // 主键ID
			CouponId:       detail.CouponId,       // 优惠券ID
			MemberId:       detail.MemberId,       // 用户ID
			GetTime:        detail.GetTime,        // 领取时间
			GetType:        detail.GetType,        // 获取类型：0->后台赠送；1->主动获取
			UseTime:        detail.UseTime,        // 使用时间
			OrderId:        detail.OrderId,        // 使用订单ID
			OrderAmount:    detail.OrderAmount,    // 订单金额
			DiscountAmount: detail.DiscountAmount, // 优惠金额
			Status:         detail.Status,         // 状态：0-未使用，1-已使用，2-已过期，3-已失效
			InvalidTime:    detail.InvalidTime,    // 失效时间
			InvalidReason:  detail.InvalidReason,  // 失效原因
			CreateTime:     detail.CreateTime,     // 创建时间
		})
	}

	return &types.ListCouponHistoryResp{
		Data:    list,
		Code:    0,
		Message: "获取会员优惠券历史列表成功",
	}, nil
}
