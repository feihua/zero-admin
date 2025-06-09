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

// QuerySeckillProductListLogic 查询秒杀商品列表
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type QuerySeckillProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySeckillProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillProductListLogic {
	return &QuerySeckillProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySeckillProductList 查询秒杀商品列表
func (l *QuerySeckillProductListLogic) QuerySeckillProductList(req *types.QuerySeckillProductListReq) (resp *types.QuerySeckillProductListResp, err error) {
	result, err := l.svcCtx.SeckillProductService.QuerySeckillProductList(l.ctx, &smsclient.QuerySeckillProductListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		ActivityId: req.ActivityId, // 活动ID
		SessionId:  req.SessionId,  // 秒杀场次ID
		Status:     req.Status,     // 状态：0-未上架，1-已上架
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字秒杀商品列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QuerySeckillProductListData

	for _, detail := range result.List {
		list = append(list, &types.QuerySeckillProductListData{
			Id:           detail.Id,                    // ID
			ActivityId:   detail.ActivityId,            // 活动ID
			SessionId:    detail.SessionId,             // 秒杀场次ID
			SkuId:        detail.SkuId,                 // 商品SKU ID
			SkuName:      detail.SkuName,               // 商品名称
			SeckillPrice: float64(detail.SeckillPrice), // 秒杀价格
			SeckillStock: detail.SeckillStock,          // 秒杀库存
			StockLocked:  detail.StockLocked,           // 锁定库存
			PerLimit:     detail.PerLimit,              // 每人限购数量
			Sort:         detail.Sort,                  // 排序
			Status:       detail.Status,                // 状态：0-未上架，1-已上架
			CreateBy:     detail.CreateBy,              // 创建人ID
			CreateTime:   detail.CreateTime,            // 创建时间
			UpdateBy:     detail.UpdateBy,              // 更新人ID
			UpdateTime:   detail.UpdateTime,            // 更新时间

		})
	}

	return &types.QuerySeckillProductListResp{
		Code:     "000000",
		Message:  "查询秒杀商品列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
