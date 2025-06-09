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

// DeleteSeckillActivityLogic 删除秒杀活动
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type DeleteSeckillActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSeckillActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSeckillActivityLogic {
	return &DeleteSeckillActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteSeckillActivity 删除秒杀活动
func (l *DeleteSeckillActivityLogic) DeleteSeckillActivity(req *types.DeleteSeckillActivityReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.SeckillActivityService.DeleteSeckillActivity(l.ctx, &smsclient.DeleteSeckillActivityReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除秒杀活动失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除秒杀活动成功",
	}, nil
}
