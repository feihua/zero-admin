package couponhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponHistoryDeleteLogic
/*
Author: LiuFeiHua
Date: 2024/5/6 11:35
*/
type CouponHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponHistoryDeleteLogic {
	return &CouponHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponHistoryDelete 删除优惠券使用记录
func (l *CouponHistoryDeleteLogic) CouponHistoryDelete(in *smsclient.CouponHistoryDeleteReq) (*smsclient.CouponHistoryDeleteResp, error) {
	q := query.SmsCouponHistory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &smsclient.CouponHistoryDeleteResp{}, nil
}
