package seckillreservationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddSeckillReservationLogic 添加秒杀预约
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type AddSeckillReservationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSeckillReservationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeckillReservationLogic {
	return &AddSeckillReservationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddSeckillReservation 添加秒杀预约
func (l *AddSeckillReservationLogic) AddSeckillReservation(in *smsclient.AddSeckillReservationReq) (*smsclient.AddSeckillReservationResp, error) {
	q := query.SmsSeckillReservation

	item := &model.SmsSeckillReservation{
		UserID:     in.UserId,     // 用户ID
		ActivityID: in.ActivityId, // 活动ID
		ProductID:  in.ProductId,  // 秒杀商品ID
		Status:     in.Status,     // 状态：0-已预约，1-已参与，2-已取消
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加秒杀预约失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加秒杀预约失败")
	}

	return &smsclient.AddSeckillReservationResp{}, nil
}
