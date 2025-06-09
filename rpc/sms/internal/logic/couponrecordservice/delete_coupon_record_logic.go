package couponrecordservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCouponRecordLogic 删除优惠券领取记录
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type DeleteCouponRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCouponRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCouponRecordLogic {
	return &DeleteCouponRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCouponRecord 删除优惠券领取记录
func (l *DeleteCouponRecordLogic) DeleteCouponRecord(in *smsclient.DeleteCouponRecordReq) (*smsclient.DeleteCouponRecordResp, error) {
	q := query.SmsCouponRecord

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除优惠券领取记录失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除优惠券领取记录失败")
	}

	return &smsclient.DeleteCouponRecordResp{}, nil
}
