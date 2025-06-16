package productskuservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// UpdateProductSkuLogic 更新商品SKU
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type UpdateProductSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSkuLogic {
	return &UpdateProductSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductSku 更新商品SKU
func (l *UpdateProductSkuLogic) UpdateProductSku(in *pmsclient.UpdateProductSkuReq) (*pmsclient.UpdateProductSkuResp, error) {
	q := query.PmsProductSku

	var skuIds []int64
	var skuStockList []*model.PmsProductSku
	for _, item := range in.Data {
		skuIds = append(skuIds, item.Id)
		elems := &model.PmsProductSku{
			ID:             item.Id,                      // 商品SpuId
			SpuID:          item.SpuId,                   // 商品SpuId
			Name:           item.Name,                    // SKU名称
			SkuCode:        item.SkuCode,                 // SKU编码
			MainPic:        item.MainPic,                 // 主图
			AlbumPics:      item.AlbumPics,               // 图片集
			Price:          float64(item.Price),          // 价格
			PromotionPrice: float64(item.PromotionPrice), // 单品促销价格
			Stock:          item.Stock,                   // 库存
			LowStock:       item.LowStock,                // 预警库存
			SpecData:       item.SpecData,                // 规格数据
			Weight:         float64(item.Weight),         // 重量(kg)
			PublishStatus:  item.PublishStatus,           // 上架状态：0-下架，1-上架
			VerifyStatus:   item.VerifyStatus,            // 审核状态：0-未审核，1-审核通过，2-审核不通过
			Sort:           item.Sort,                    // 排序
			UpdateBy:       &item.UpdateBy,               // 更新人ID
		}

		if len(item.PromotionStartTime) > 0 {
			startTime, _ := time.Parse("2006-01-02 15:04:05", item.PromotionStartTime)
			endTime, _ := time.Parse("2006-01-02 15:04:05", item.PromotionEndTime)
			elems.PromotionStartTime = &startTime // 促销开始时间
			elems.PromotionEndTime = &endTime     // 促销结束时间
		}
		skuStockList = append(skuStockList, elems)
	}
	// 1.先删除
	_, err := q.WithContext(l.ctx).Where(q.ID.In(skuIds...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "更新sku的库存失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新sku的库存失败")
	}

	// 2.后添加
	err = q.WithContext(l.ctx).CreateInBatches(skuStockList, len(skuStockList))

	if err != nil {
		logc.Errorf(l.ctx, "更新sku的库存失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新sku的库存失败,错误信息:%s", err.Error()))
	}

	return &pmsclient.UpdateProductSkuResp{}, nil
}
