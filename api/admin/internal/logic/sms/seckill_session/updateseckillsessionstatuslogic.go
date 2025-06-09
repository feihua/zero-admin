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

// UpdateSeckillSessionStatusLogic 更新秒杀场次状态状态
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:04
*/
type UpdateSeckillSessionStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSeckillSessionStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillSessionStatusLogic {
	return &UpdateSeckillSessionStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateSeckillSessionStatus 更新秒杀场次状态
func (l *UpdateSeckillSessionStatusLogic) UpdateSeckillSessionStatus(req *types.UpdateSeckillSessionStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.SeckillSessionService.UpdateSeckillSessionStatus(l.ctx, &smsclient.UpdateSeckillSessionStatusReq{
		Ids:      req.Ids,    // 秒杀场次ID
		Status:   req.Status, // 状态：0-禁用，1-启用
		UpdateBy: userId,     // 更新人ID

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀场次状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新秒杀场次状态成功",
	}, nil
}
