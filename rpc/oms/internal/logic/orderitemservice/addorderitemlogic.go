package orderitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderItemLogic 添加订单中所包含的商品
/*
Author: LiuFeiHua
Date: 2024/6/12 10:12
*/
type AddOrderItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderItemLogic {
	return &AddOrderItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderItem 添加订单中所包含的商品
func (l *AddOrderItemLogic) AddOrderItem(in *omsclient.AddOrderItemReq) (*omsclient.AddOrderItemResp, error) {
	err := query.OmsOrderItem.WithContext(l.ctx).Create(&model.OmsOrderItem{
		OrderID:           in.OrderId,           // 订单id
		OrderSn:           in.OrderSn,           // 订单编号
		ProductID:         in.ProductId,         // 商品id
		ProductPic:        in.ProductPic,        // 商品图片
		ProductName:       in.ProductName,       // 商品名称
		ProductBrand:      in.ProductBrand,      // 商品品牌
		ProductSn:         in.ProductSn,         // 货号
		ProductPrice:      in.ProductPrice,      // 销售价格
		ProductQuantity:   in.ProductQuantity,   // 购买数量
		ProductSkuID:      in.ProductSkuId,      // 商品sku编号
		ProductSkuCode:    in.ProductSkuCode,    // 商品sku条码
		ProductCategoryID: in.ProductCategoryId, // 商品分类id
		PromotionName:     in.PromotionName,     // 商品促销名称
		PromotionAmount:   in.PromotionAmount,   // 商品促销分解金额
		CouponAmount:      in.CouponAmount,      // 优惠券优惠分解金额
		IntegrationAmount: in.IntegrationAmount, // 积分优惠分解金额
		RealAmount:        in.RealAmount,        // 该商品经过优惠后的分解金额
		GiftIntegration:   in.GiftIntegration,   //
		GiftGrowth:        in.GiftGrowth,        //
		ProductAttr:       in.ProductAttr,       // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.AddOrderItemResp{}, nil
}
