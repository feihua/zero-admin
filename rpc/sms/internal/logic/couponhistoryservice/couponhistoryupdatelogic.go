package couponhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryUpdateLogic {
	return &CouponHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponHistoryUpdateLogic) CouponHistoryUpdate(in *smsclient.CouponHistoryUpdateReq) (*smsclient.CouponHistoryUpdateResp, error) {
	q := query.SmsCouponHistory
	_, err := q.WithContext(l.ctx).Updates(&model.SmsCouponHistory{
		ID:             in.Id,
		CouponID:       in.CouponId,
		MemberID:       in.MemberId,
		CouponCode:     in.CouponCode,
		MemberNickname: in.MemberNickname,
		GetType:        in.GetType,
		CreateTime:     time.Now(),
		UseStatus:      in.UseStatus,
		UseTime:        time.Now(),
		OrderID:        in.OrderId,
		OrderSn:        in.OrderSn,
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.CouponHistoryUpdateResp{}, nil
}
