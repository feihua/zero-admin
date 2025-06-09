package seckill_product

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

// DeleteSeckillProductLogic 删除秒杀商品
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type DeleteSeckillProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSeckillProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSeckillProductLogic {
	return &DeleteSeckillProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteSeckillProduct 删除秒杀商品
func (l *DeleteSeckillProductLogic) DeleteSeckillProduct(req *types.DeleteSeckillProductReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.SeckillProductService.DeleteSeckillProduct(l.ctx, &smsclient.DeleteSeckillProductReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除秒杀商品失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除秒杀商品成功",
	}, nil
}
