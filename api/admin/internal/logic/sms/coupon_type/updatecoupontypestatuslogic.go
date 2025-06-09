package coupon_type

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

// UpdateCouponTypeStatusLogic 更新优惠券类型状态状态
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type UpdateCouponTypeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCouponTypeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponTypeStatusLogic {
	return &UpdateCouponTypeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCouponTypeStatus 更新优惠券类型状态
func (l *UpdateCouponTypeStatusLogic) UpdateCouponTypeStatus(req *types.UpdateCouponTypeStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CouponTypeService.UpdateCouponTypeStatus(l.ctx, &smsclient.UpdateCouponTypeStatusReq{
		Ids:      req.Ids,    // 主键ID
		Status:   req.Status, // 是否启用
		UpdateBy: userId,     // 更新人ID

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券类型状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新优惠券类型状态成功",
	}, nil
}
