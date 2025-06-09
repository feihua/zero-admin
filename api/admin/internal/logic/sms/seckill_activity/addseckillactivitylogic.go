package seckill_activity

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddSeckillActivityLogic 添加秒杀活动
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type AddSeckillActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSeckillActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeckillActivityLogic {
	return &AddSeckillActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddSeckillActivity 添加秒杀活动
func (l *AddSeckillActivityLogic) AddSeckillActivity(req *types.AddSeckillActivityReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.SeckillActivityService.AddSeckillActivity(l.ctx, &smsclient.AddSeckillActivityReq{
		Name:        req.Name,        // 活动名称
		Description: req.Description, // 活动描述
		StartTime:   req.StartTime,   // 开始时间
		EndTime:     req.EndTime,     // 结束时间
		Status:      req.Status,      // 活动状态：0-未开始，1-进行中，2-已结束，3-已取消
		IsEnabled:   req.IsEnabled,   // 是否启用
		CreateBy:    userId,          // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加秒杀活动失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "添加秒杀活动成功",
	}, nil
}
