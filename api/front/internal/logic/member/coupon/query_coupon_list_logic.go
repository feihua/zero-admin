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

// QueryCouponListLogic 获取用户优惠券列表
/*
Author: LiuFeiHua
Date: 2025/6/19 11:15
*/
type QueryCouponListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponListLogic {
	return &QueryCouponListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCouponList 获取用户优惠券列表
func (l *QueryCouponListLogic) QueryCouponList(req *types.ListCouponReq) (resp *types.ListCouponResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	couponList, err := l.svcCtx.CouponRecordService.QueryMemberCouponList(l.ctx, &smsclient.QueryMemberCouponListReq{
		MemberId: memberId,
		Status:   req.UseStatus,
	})
	if err != nil {
		logc.Errorf(l.ctx, "获取用户优惠券列表失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	list := make([]*types.CouponData, 0)
	for _, detail := range couponList.List {
		list = append(list, &types.CouponData{
			Id:          detail.Id,          // 优惠券ID
			TypeId:      detail.TypeId,      // 优惠券类型ID
			Name:        detail.Name,        // 优惠券名称
			Code:        detail.Code,        // 优惠券码
			Amount:      detail.Amount,      // 优惠金额/折扣率
			MinAmount:   detail.MinAmount,   // 最低使用金额
			StartTime:   detail.StartTime,   // 生效时间
			EndTime:     detail.EndTime,     // 失效时间
			PerLimit:    detail.PerLimit,    // 每人限领数量
			Status:      detail.Status,      // 状态：0-未开始，1-进行中，2-已结束，3-已取消
			Description: detail.Description, // 使用说明
			ScopeType:   detail.ScopeType,   // 范围类型：0-全场通用，1-指定分类，2-指定商品
		})
	}

	return &types.ListCouponResp{
		Data:    list,
		Code:    0,
		Message: "查询会员优惠券成功",
	}, nil
}
