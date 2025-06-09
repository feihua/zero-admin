package coupon_type

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

// DeleteCouponTypeLogic 删除优惠券类型
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type DeleteCouponTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCouponTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCouponTypeLogic {
	return &DeleteCouponTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteCouponType 删除优惠券类型
func (l *DeleteCouponTypeLogic) DeleteCouponType(req *types.DeleteCouponTypeReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.CouponTypeService.DeleteCouponType(l.ctx, &smsclient.DeleteCouponTypeReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除优惠券类型失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除优惠券类型成功",
	}, nil
}
