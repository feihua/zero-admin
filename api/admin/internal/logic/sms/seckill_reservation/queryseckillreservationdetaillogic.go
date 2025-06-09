package seckill_reservation

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySeckillReservationDetailLogic 查询秒杀预约详情
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:04
*/
type QuerySeckillReservationDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySeckillReservationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillReservationDetailLogic {
	return &QuerySeckillReservationDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySeckillReservationDetail 查询秒杀预约详情
func (l *QuerySeckillReservationDetailLogic) QuerySeckillReservationDetail(req *types.QuerySeckillReservationDetailReq) (resp *types.QuerySeckillReservationDetailResp, err error) {

	detail, err := l.svcCtx.SeckillReservationService.QuerySeckillReservationDetail(l.ctx, &smsclient.QuerySeckillReservationDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询秒杀预约详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QuerySeckillReservationDetailData{
		Id:         detail.Id,         // ID
		UserId:     detail.UserId,     // 用户ID
		ActivityId: detail.ActivityId, // 活动ID
		ProductId:  detail.ProductId,  // 秒杀商品ID
		Status:     detail.Status,     // 状态：0-已预约，1-已参与，2-已取消
		CreateTime: detail.CreateTime, //
		UpdateTime: detail.UpdateTime, //

	}
	return &types.QuerySeckillReservationDetailResp{
		Code:    "000000",
		Message: "查询秒杀预约成功",
		Data:    data,
	}, nil
}
