package seckill_activity

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

// QuerySeckillActivityListLogic 查询秒杀活动列表
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type QuerySeckillActivityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySeckillActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillActivityListLogic {
	return &QuerySeckillActivityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySeckillActivityList 查询秒杀活动列表
func (l *QuerySeckillActivityListLogic) QuerySeckillActivityList(req *types.QuerySeckillActivityListReq) (resp *types.QuerySeckillActivityListResp, err error) {
	result, err := l.svcCtx.SeckillActivityService.QuerySeckillActivityList(l.ctx, &smsclient.QuerySeckillActivityListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		Name:      req.Name,      // 活动名称
		StartTime: req.StartTime, // 开始时间
		EndTime:   req.EndTime,   // 结束时间
		Status:    req.Status,    // 活动状态：0-未开始，1-进行中，2-已结束，3-已取消
		IsEnabled: req.IsEnabled, // 是否启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字秒杀活动列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QuerySeckillActivityListData

	for _, detail := range result.List {
		list = append(list, &types.QuerySeckillActivityListData{
			Id:          detail.Id,          // 编号
			Name:        detail.Name,        // 活动名称
			Description: detail.Description, // 活动描述
			StartTime:   detail.StartTime,   // 开始时间
			EndTime:     detail.EndTime,     // 结束时间
			Status:      detail.Status,      // 活动状态：0-未开始，1-进行中，2-已结束，3-已取消
			IsEnabled:   detail.IsEnabled,   // 是否启用
			CreateBy:    detail.CreateBy,    // 创建人ID
			CreateTime:  detail.CreateTime,  // 创建时间
			UpdateBy:    detail.UpdateBy,    // 更新人ID
			UpdateTime:  detail.UpdateTime,  // 更新时间

		})
	}

	return &types.QuerySeckillActivityListResp{
		Code:     "000000",
		Message:  "查询秒杀活动列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
