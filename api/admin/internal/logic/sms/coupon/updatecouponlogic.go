package coupon

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCouponLogic 更新优惠券
/*
Author: LiuFeiHua
Date: 2025/06/12 10:40:14
*/
type UpdateCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponLogic {
	return &UpdateCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCoupon 更新优惠券
func (l *UpdateCouponLogic) UpdateCoupon(req *types.UpdateCouponReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	couponReq := &smsclient.UpdateCouponReq{
		Id:            req.Id,            // 优惠券ID
		TypeId:        req.TypeId,        // 优惠券类型ID
		Name:          req.Name,          // 优惠券名称
		Code:          req.Code,          // 优惠券码
		Amount:        req.Amount,        // 优惠金额/折扣率
		MinAmount:     req.MinAmount,     // 最低使用金额
		StartTime:     req.StartTime,     // 生效时间
		EndTime:       req.EndTime,       // 失效时间
		TotalCount:    req.TotalCount,    // 发放总量
		ReceivedCount: req.ReceivedCount, // 已领取数量
		UsedCount:     req.UsedCount,     // 已使用数量
		PerLimit:      req.PerLimit,      // 每人限领数量
		Status:        req.Status,        // 状态：0-未开始，1-进行中，2-已结束，3-已取消
		IsEnabled:     req.IsEnabled,     // 是否启用
		Description:   req.Description,   // 使用说明
		UpdateBy:      userId,            // 更新人ID
	}

	if len(req.ScopeData) > 0 {
		var list []*smsclient.CouponScopeData

		for _, detail := range req.ScopeData {
			list = append(list, &smsclient.CouponScopeData{
				ScopeType: detail.ScopeType,
				ScopeId:   detail.ScopeId,
			})
		}
		couponReq.Scopes = list
	}

	_, err = l.svcCtx.CouponService.UpdateCoupon(l.ctx, couponReq)

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新优惠券成功",
	}, nil
}
