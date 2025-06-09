package seckillproductservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddSeckillProductLogic 添加秒杀商品
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type AddSeckillProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSeckillProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeckillProductLogic {
	return &AddSeckillProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddSeckillProduct 添加秒杀商品
func (l *AddSeckillProductLogic) AddSeckillProduct(in *smsclient.AddSeckillProductReq) (*smsclient.AddSeckillProductResp, error) {
	q := query.SmsSeckillProduct

	var data []*model.SmsSeckillProduct
	for _, x := range in.Data {
		data = append(data, &model.SmsSeckillProduct{
			ActivityID:   x.ActivityId,            // 活动ID
			SessionID:    x.SessionId,             // 秒杀场次ID
			SkuID:        x.SkuId,                 // 商品SKU ID
			SkuName:      x.SkuName,               // 商品名称
			SeckillPrice: float64(x.SeckillPrice), // 秒杀价格
			SeckillStock: x.SeckillStock,          // 秒杀库存
			StockLocked:  x.StockLocked,           // 锁定库存
			PerLimit:     x.PerLimit,              // 每人限购数量
			Sort:         x.Sort,                  // 排序
			Status:       x.Status,                // 状态：0-未上架，1-已上架
			CreateBy:     x.CreateBy,              // 创建人ID
		})
	}

	err := q.WithContext(l.ctx).CreateInBatches(data, len(data))
	if err != nil {
		logc.Errorf(l.ctx, "添加秒杀商品失败,参数:%+v,异常:%s", data, err.Error())
		return nil, errors.New("添加秒杀商品失败")
	}

	return &smsclient.AddSeckillProductResp{}, nil
}
