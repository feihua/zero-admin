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
	"gorm.io/gorm"
	"time"
)

// UpdateSeckillProductLogic 更新秒杀商品
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:51
*/
type UpdateSeckillProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSeckillProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillProductLogic {
	return &UpdateSeckillProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSeckillProduct 更新秒杀商品
func (l *UpdateSeckillProductLogic) UpdateSeckillProduct(in *smsclient.UpdateSeckillProductReq) (*smsclient.UpdateSeckillProductResp, error) {
	q := query.SmsSeckillProduct.WithContext(l.ctx)

	// 1.根据秒杀商品id查询秒杀商品是否已存在
	detail, err := q.Where(query.SmsSeckillProduct.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "秒杀商品不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("秒杀商品不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询秒杀商品异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询秒杀商品异常")
	}

	now := time.Now()
	item := &model.SmsSeckillProduct{
		ID:           in.Id,                    // ID
		ActivityID:   in.ActivityId,            // 活动ID
		SessionID:    in.SessionId,             // 秒杀场次ID
		SkuID:        in.SkuId,                 // 商品SKU ID
		SkuName:      in.SkuName,               // 商品名称
		SeckillPrice: float64(in.SeckillPrice), // 秒杀价格
		SeckillStock: in.SeckillStock,          // 秒杀库存
		StockLocked:  in.StockLocked,           // 锁定库存
		PerLimit:     in.PerLimit,              // 每人限购数量
		Sort:         in.Sort,                  // 排序
		Status:       in.Status,                // 状态：0-未上架，1-已上架
		CreateBy:     detail.CreateBy,          // 创建人ID
		CreateTime:   detail.CreateTime,        // 创建时间
		UpdateBy:     &in.UpdateBy,             // 更新人ID
		UpdateTime:   &now,                     // 更新时间
	}

	// 2.秒杀商品存在时,则直接更新秒杀商品
	err = l.svcCtx.DB.Model(&model.SmsSeckillProduct{}).WithContext(l.ctx).Where(query.SmsSeckillProduct.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀商品失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新秒杀商品失败")
	}

	return &smsclient.UpdateSeckillProductResp{}, nil
}
