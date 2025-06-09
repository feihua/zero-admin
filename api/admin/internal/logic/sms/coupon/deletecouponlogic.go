package coupon

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

// DeleteCouponLogic 删除优惠券
/*
Author: LiuFeiHua
Date: 2025/06/12 10:40:14
*/
type DeleteCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCouponLogic {
	return &DeleteCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteCoupon 删除优惠券
func (l *DeleteCouponLogic) DeleteCoupon(req *types.DeleteCouponReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.CouponService.DeleteCoupon(l.ctx, &smsclient.DeleteCouponReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除优惠券失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除优惠券成功",
	}, nil
}
