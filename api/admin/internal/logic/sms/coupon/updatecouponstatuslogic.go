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

// UpdateCouponStatusLogic 更新优惠券状态状态
/*
Author: LiuFeiHua
Date: 2025/06/12 10:40:14
*/
type UpdateCouponStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCouponStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponStatusLogic {
	return &UpdateCouponStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCouponStatus 更新优惠券状态
func (l *UpdateCouponStatusLogic) UpdateCouponStatus(req *types.UpdateSmsCouponStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CouponService.UpdateCouponStatus(l.ctx, &smsclient.UpdateCouponStatusReq{
		Ids:      req.Ids,    // 优惠券ID
		Status:   req.Status, // 状态：0-未开始，1-进行中，2-已结束，3-已取消
		UpdateBy: userId,     // 更新人ID

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新优惠券状态成功",
	}, nil
}
