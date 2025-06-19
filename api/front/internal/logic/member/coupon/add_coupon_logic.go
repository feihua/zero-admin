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

// AddCouponLogic 领取指定优惠券
/*
Author: LiuFeiHua
Date: 2025/6/19 11:13
*/
type AddCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponLogic {
	return &AddCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddCoupon 领取指定优惠券
func (l *AddCouponLogic) AddCoupon(req *types.AddCouponReq) (resp *types.AddCouponResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CouponRecordService.AddCouponRecord(l.ctx, &smsclient.AddCouponRecordReq{
		CouponId: req.CouponId,
		MemberId: memberId,
		GetType:  1, // 主动领取
	})

	if err != nil {
		logc.Errorf(l.ctx, "领取指定优惠券失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddCouponResp{
		Code:    0,
		Message: "领取优惠券成功",
	}, nil
}
