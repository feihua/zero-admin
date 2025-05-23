package member_consume_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCouponStatusLogic 更新是否可以和优惠券同用
/*
Author: LiuFeiHua
Date: 2025/5/23 14:14
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

// UpdateCouponStatus 更新是否可以和优惠券同用
func (l *UpdateCouponStatusLogic) UpdateCouponStatus(req *types.UpdateCouponStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberConsumeSettingService.UpdateCouponStatus(l.ctx, &umsclient.UpdateCouponStatusReq{
		Id:           req.Id,           //
		CouponStatus: req.CouponStatus, // 是否可以和优惠券同用；0->不可以；1->可以
		UpdateBy:     userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新积分消费设置状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新积分消费设置状态成功",
	}, nil
}
