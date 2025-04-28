package coupon

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponDeleteLogic 优惠券
/*
Author: LiuFeiHua
Date: 2024/5/15 10:09
*/
type CouponDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponDeleteLogic {
	return CouponDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CouponDelete 删除优惠券
func (l *CouponDeleteLogic) CouponDelete(req *types.DeleteCouponReq) (*types.BaseResp, error) {
	_, err := l.svcCtx.CouponService.DeleteCoupon(l.ctx, &smsclient.DeleteCouponReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除优惠券异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	return res.Success()
}
