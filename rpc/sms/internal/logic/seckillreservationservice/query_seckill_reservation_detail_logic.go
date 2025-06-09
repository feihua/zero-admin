package seckillreservationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QuerySeckillReservationDetailLogic 查询秒杀预约详情
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type QuerySeckillReservationDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySeckillReservationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillReservationDetailLogic {
	return &QuerySeckillReservationDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySeckillReservationDetail 查询秒杀预约详情
func (l *QuerySeckillReservationDetailLogic) QuerySeckillReservationDetail(in *smsclient.QuerySeckillReservationDetailReq) (*smsclient.QuerySeckillReservationDetailResp, error) {
	item, err := query.SmsSeckillReservation.WithContext(l.ctx).Where(query.SmsSeckillReservation.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "秒杀预约不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("秒杀预约不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询秒杀预约异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询秒杀预约异常")
	}

	data := &smsclient.QuerySeckillReservationDetailResp{
		Id:         item.ID,                                 // ID
		UserId:     item.UserID,                             // 用户ID
		ActivityId: item.ActivityID,                         // 活动ID
		ProductId:  item.ProductID,                          // 秒杀商品ID
		Status:     item.Status,                             // 状态：0-已预约，1-已参与，2-已取消
		CreateTime: time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	return data, nil
}
