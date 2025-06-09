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
	"gorm.io/gorm"
	"time"
)

// UpdateSeckillReservationLogic 更新秒杀预约
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type UpdateSeckillReservationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSeckillReservationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillReservationLogic {
	return &UpdateSeckillReservationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSeckillReservation 更新秒杀预约
func (l *UpdateSeckillReservationLogic) UpdateSeckillReservation(in *smsclient.UpdateSeckillReservationReq) (*smsclient.UpdateSeckillReservationResp, error) {
	q := query.SmsSeckillReservation.WithContext(l.ctx)

	// 1.根据秒杀预约id查询秒杀预约是否已存在
	detail, err := q.Where(query.SmsSeckillReservation.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "秒杀预约不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("秒杀预约不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询秒杀预约异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询秒杀预约异常")
	}

	now := time.Now()
	item := &model.SmsSeckillReservation{
		ID:         in.Id,             // ID
		UserID:     in.UserId,         // 用户ID
		ActivityID: in.ActivityId,     // 活动ID
		ProductID:  in.ProductId,      // 秒杀商品ID
		Status:     in.Status,         // 状态：0-已预约，1-已参与，2-已取消
		CreateTime: detail.CreateTime, // 创建时间
		UpdateTime: &now,              // 更新时间
	}

	// 2.秒杀预约存在时,则直接更新秒杀预约
	err = l.svcCtx.DB.Model(&model.SmsSeckillReservation{}).WithContext(l.ctx).Where(query.SmsSeckillReservation.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀预约失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新秒杀预约失败")
	}

	return &smsclient.UpdateSeckillReservationResp{}, nil
}
