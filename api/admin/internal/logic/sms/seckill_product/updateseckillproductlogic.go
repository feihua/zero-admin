package seckill_product

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

// UpdateSeckillProductLogic 更新秒杀商品
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type UpdateSeckillProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSeckillProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillProductLogic {
	return &UpdateSeckillProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateSeckillProduct 更新秒杀商品
func (l *UpdateSeckillProductLogic) UpdateSeckillProduct(req *types.UpdateSeckillProductReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.SeckillProductService.UpdateSeckillProduct(l.ctx, &smsclient.UpdateSeckillProductReq{
		Id:           req.Id,           // ID
		ActivityId:   req.ActivityId,   // 活动ID
		SessionId:    req.SessionId,    // 秒杀场次ID
		SkuId:        req.SkuId,        // 商品SKU ID
		SkuName:      req.SkuName,      // 商品名称
		SeckillPrice: req.SeckillPrice, // 秒杀价格
		SeckillStock: req.SeckillStock, // 秒杀库存
		StockLocked:  req.StockLocked,  // 锁定库存
		PerLimit:     req.PerLimit,     // 每人限购数量
		Sort:         req.Sort,         // 排序
		Status:       req.Status,       // 状态：0-未上架，1-已上架
		UpdateBy:     userId,           // 更新人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀商品失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新秒杀商品成功",
	}, nil
}
