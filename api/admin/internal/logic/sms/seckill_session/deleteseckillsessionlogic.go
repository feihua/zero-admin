package seckill_session

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

// DeleteSeckillSessionLogic 删除秒杀场次
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:04
*/
type DeleteSeckillSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSeckillSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSeckillSessionLogic {
	return &DeleteSeckillSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteSeckillSession 删除秒杀场次
func (l *DeleteSeckillSessionLogic) DeleteSeckillSession(req *types.DeleteSeckillSessionReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.SeckillSessionService.DeleteSeckillSession(l.ctx, &smsclient.DeleteSeckillSessionReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除秒杀场次失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除秒杀场次成功",
	}, nil
}
