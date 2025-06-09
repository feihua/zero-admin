package coupon_scope

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCouponScopeLogic 更新优惠券使用范围
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type UpdateCouponScopeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCouponScopeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCouponScopeLogic {
	return &UpdateCouponScopeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCouponScope 更新优惠券使用范围
func (l *UpdateCouponScopeLogic) UpdateCouponScope(req *types.UpdateCouponScopeReq) (resp *types.BaseResp, err error) {

	_, err = l.svcCtx.CouponScopeService.UpdateCouponScope(l.ctx, &smsclient.UpdateCouponScopeReq{
		Id:        req.Id,        // 主键ID
		CouponId:  req.CouponId,  // 优惠券ID
		ScopeType: req.ScopeType, // 范围类型：0-全场，1-分类，2-商品
		ScopeId:   req.ScopeId,   // 范围ID（分类ID或商品ID）
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新优惠券使用范围失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新优惠券使用范围成功",
	}, nil
}
