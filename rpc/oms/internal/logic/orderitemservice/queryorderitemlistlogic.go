package orderitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderItemListLogic 查询订单中所包含的商品列表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:13
*/
type QueryOrderItemListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderItemListLogic {
	return &QueryOrderItemListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderItemList 查询订单中所包含的商品列表
func (l *QueryOrderItemListLogic) QueryOrderItemList(in *omsclient.QueryOrderItemListReq) (*omsclient.QueryOrderItemListResp, error) {
	q := query.OmsOrderItem.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询订单详情列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*omsclient.OrderItemListData
	for _, item := range result {

		list = append(list, &omsclient.OrderItemListData{
			Id:                item.ID,                //
			OrderId:           item.OrderID,           // 订单id
			OrderSn:           item.OrderSn,           // 订单编号
			ProductId:         item.ProductID,         // 商品id
			ProductPic:        item.ProductPic,        // 商品图片
			ProductName:       item.ProductName,       // 商品名称
			ProductBrand:      item.ProductBrand,      // 商品品牌
			ProductSn:         item.ProductSn,         // 货号
			ProductPrice:      item.ProductPrice,      // 销售价格
			ProductQuantity:   item.ProductQuantity,   // 购买数量
			ProductSkuId:      item.ProductSkuID,      // 商品sku编号
			ProductSkuCode:    item.ProductSkuCode,    // 商品sku条码
			ProductCategoryId: item.ProductCategoryID, // 商品分类id
			PromotionName:     item.PromotionName,     // 商品促销名称
			PromotionAmount:   item.PromotionAmount,   // 商品促销分解金额
			CouponAmount:      item.CouponAmount,      // 优惠券优惠分解金额
			IntegrationAmount: item.IntegrationAmount, // 积分优惠分解金额
			RealAmount:        item.RealAmount,        // 该商品经过优惠后的分解金额
			GiftIntegration:   item.GiftIntegration,   //
			GiftGrowth:        item.GiftGrowth,        //
			ProductAttr:       item.ProductAttr,       // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
		})
	}

	return &omsclient.QueryOrderItemListResp{
		Total: count,
		List:  list,
	}, nil

}
