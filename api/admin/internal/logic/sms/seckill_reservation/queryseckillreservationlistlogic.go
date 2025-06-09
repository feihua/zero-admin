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

// QuerySeckillReservationListLogic 查询秒杀预约列表
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:04
*/
type QuerySeckillReservationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySeckillReservationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillReservationListLogic {
	return &QuerySeckillReservationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySeckillReservationList 查询秒杀预约列表
func (l *QuerySeckillReservationListLogic) QuerySeckillReservationList(req *types.QuerySeckillReservationListReq) (resp *types.QuerySeckillReservationListResp, err error) {
	result, err := l.svcCtx.SeckillReservationService.QuerySeckillReservationList(l.ctx, &smsclient.QuerySeckillReservationListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		UserId:     req.UserId,     // 用户ID
		ActivityId: req.ActivityId, // 活动ID
		ProductId:  req.ProductId,  // 秒杀商品ID
		Status:     req.Status,     // 状态：0-已预约，1-已参与，2-已取消
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字秒杀预约列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QuerySeckillReservationListData

	for _, detail := range result.List {
		list = append(list, &types.QuerySeckillReservationListData{
			Id:         detail.Id,         // ID
			UserId:     detail.UserId,     // 用户ID
			ActivityId: detail.ActivityId, // 活动ID
			ProductId:  detail.ProductId,  // 秒杀商品ID
			Status:     detail.Status,     // 状态：0-已预约，1-已参与，2-已取消
			CreateTime: detail.CreateTime, //
			UpdateTime: detail.UpdateTime, //

		})
	}

	return &types.QuerySeckillReservationListResp{
		Code:     "000000",
		Message:  "查询秒杀预约列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
