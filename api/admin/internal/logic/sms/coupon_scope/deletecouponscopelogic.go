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

// DeleteCouponScopeLogic 删除优惠券使用范围
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type DeleteCouponScopeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCouponScopeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCouponScopeLogic {
	return &DeleteCouponScopeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteCouponScope 删除优惠券使用范围
func (l *DeleteCouponScopeLogic) DeleteCouponScope(req *types.DeleteCouponScopeReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.CouponScopeService.DeleteCouponScope(l.ctx, &smsclient.DeleteCouponScopeReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除优惠券使用范围失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除优惠券使用范围成功",
	}, nil
}
