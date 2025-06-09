package seckillreservationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSeckillReservationStatusLogic 更新秒杀预约
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type UpdateSeckillReservationStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSeckillReservationStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillReservationStatusLogic {
	return &UpdateSeckillReservationStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSeckillReservationStatus 更新秒杀预约状态
func (l *UpdateSeckillReservationStatusLogic) UpdateSeckillReservationStatus(in *smsclient.UpdateSeckillReservationStatusReq) (*smsclient.UpdateSeckillReservationStatusResp, error) {
	q := query.SmsSeckillReservation

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀预约状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新秒杀预约状态失败")
	}

	return &smsclient.UpdateSeckillReservationStatusResp{}, nil
}
