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

// AddSeckillProductLogic 添加秒杀商品
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type AddSeckillProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSeckillProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeckillProductLogic {
	return &AddSeckillProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddSeckillProduct 添加秒杀商品
func (l *AddSeckillProductLogic) AddSeckillProduct(req *types.AddSeckillProductReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	var list []*smsclient.AddSeckillProductData
	for _, item := range req.Data {
		list = append(list, &smsclient.AddSeckillProductData{
			ActivityId:   item.ActivityId,   // 活动ID
			SessionId:    item.SessionId,    // 秒杀场次ID
			SkuId:        item.SkuId,        // 商品SKU ID
			SkuName:      item.SkuName,      // 商品名称
			SeckillPrice: item.SeckillPrice, // 秒杀价格
			SeckillStock: item.SeckillStock, // 秒杀库存
			StockLocked:  item.StockLocked,  // 锁定库存
			PerLimit:     item.PerLimit,     // 每人限购数量
			Sort:         item.Sort,         // 排序
			Status:       item.Status,       // 状态：0-未上架，1-已上架
			CreateBy:     userId,            // 创建人ID
		})
	}
	_, err = l.svcCtx.SeckillProductService.AddSeckillProduct(l.ctx, &smsclient.AddSeckillProductReq{
		Data: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加秒杀商品失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "添加秒杀商品成功",
	}, nil
}
