package coupon

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddCouponLogic 领取指定优惠券
/*
Author: LiuFeiHua
Date: 2024/5/16 16:01
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
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, err = l.svcCtx.CouponHistoryService.CouponHistoryAdd(l.ctx, &smsclient.CouponHistoryAddReq{
		CouponId:       req.CouponId,
		MemberId:       memberId,
		CouponCode:     "哈哈还没有弄!",
		MemberNickname: l.ctx.Value("memberName").(string),
		GetType:        1, //主动领取
		UseStatus:      0, //未使用
	})

	if err != nil {
		logc.Errorf(l.ctx, "领取优惠券失败,参数：%+v,响应：%s", req, err.Error())
		return &types.AddCouponResp{
			Code:    0,
			Message: "领取优惠券失败",
		}, nil
	}

	return &types.AddCouponResp{
		Code:    0,
		Message: "领取优惠券成功",
	}, nil
}
