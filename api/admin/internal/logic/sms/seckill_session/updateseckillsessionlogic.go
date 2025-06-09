package seckill_session

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

// UpdateSeckillSessionLogic 更新秒杀场次
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:04
*/
type UpdateSeckillSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSeckillSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillSessionLogic {
	return &UpdateSeckillSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateSeckillSession 更新秒杀场次
func (l *UpdateSeckillSessionLogic) UpdateSeckillSession(req *types.UpdateSeckillSessionReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.SeckillSessionService.UpdateSeckillSession(l.ctx, &smsclient.UpdateSeckillSessionReq{
		Id:        req.Id,        // 秒杀场次ID
		Name:      req.Name,      // 场次名称
		StartTime: req.StartTime, // 开始时间
		EndTime:   req.EndTime,   // 结束时间
		Status:    req.Status,    // 状态：0-禁用，1-启用
		Sort:      req.Sort,      // 排序
		UpdateBy:  userId,        // 更新人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀场次失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新秒杀场次成功",
	}, nil
}
