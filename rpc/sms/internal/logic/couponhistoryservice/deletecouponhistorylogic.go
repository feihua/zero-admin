package couponhistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCouponHistoryLogic 删除优惠券使用、领取历史表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:28
*/
type DeleteCouponHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCouponHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCouponHistoryLogic {
	return &DeleteCouponHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCouponHistory 删除优惠券使用、领取历史表
func (l *DeleteCouponHistoryLogic) DeleteCouponHistory(in *smsclient.DeleteCouponHistoryReq) (*smsclient.DeleteCouponHistoryResp, error) {
	q := query.SmsCouponHistory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除优惠券使用失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除优惠券使用失败")
	}

	return &smsclient.DeleteCouponHistoryResp{}, nil
}
