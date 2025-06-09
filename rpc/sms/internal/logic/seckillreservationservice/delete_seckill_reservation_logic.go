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

// DeleteSeckillReservationLogic 删除秒杀预约
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type DeleteSeckillReservationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSeckillReservationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSeckillReservationLogic {
	return &DeleteSeckillReservationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteSeckillReservation 删除秒杀预约
func (l *DeleteSeckillReservationLogic) DeleteSeckillReservation(in *smsclient.DeleteSeckillReservationReq) (*smsclient.DeleteSeckillReservationResp, error) {
	q := query.SmsSeckillReservation

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除秒杀预约失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除秒杀预约失败")
	}

	return &smsclient.DeleteSeckillReservationResp{}, nil
}
