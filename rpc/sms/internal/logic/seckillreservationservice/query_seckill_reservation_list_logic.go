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
)

// QuerySeckillReservationListLogic 查询秒杀预约列表
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type QuerySeckillReservationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySeckillReservationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillReservationListLogic {
	return &QuerySeckillReservationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySeckillReservationList 查询秒杀预约列表
func (l *QuerySeckillReservationListLogic) QuerySeckillReservationList(in *smsclient.QuerySeckillReservationListReq) (*smsclient.QuerySeckillReservationListResp, error) {
	seckillReservation := query.SmsSeckillReservation
	q := seckillReservation.WithContext(l.ctx)
	if in.UserId != 2 {
		q = q.Where(seckillReservation.UserID.Eq(in.UserId))
	}
	if in.ActivityId != 2 {
		q = q.Where(seckillReservation.ActivityID.Eq(in.ActivityId))
	}
	if in.ProductId != 2 {
		q = q.Where(seckillReservation.ProductID.Eq(in.ProductId))
	}
	if in.Status != 2 {
		q = q.Where(seckillReservation.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询秒杀预约列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询秒杀预约列表失败")
	}

	var list []*smsclient.SeckillReservationListData

	for _, item := range result {
		list = append(list, &smsclient.SeckillReservationListData{
			Id:         item.ID,                                 // ID
			UserId:     item.UserID,                             // 用户ID
			ActivityId: item.ActivityID,                         // 活动ID
			ProductId:  item.ProductID,                          // 秒杀商品ID
			Status:     item.Status,                             // 状态：0-已预约，1-已参与，2-已取消
			CreateTime: time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	return &smsclient.QuerySeckillReservationListResp{
		Total: count,
		List:  list,
	}, nil
}
