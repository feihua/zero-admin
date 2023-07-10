package orderitemservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderItemListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderItemListLogic {
	return &OrderItemListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderItemListLogic) OrderItemList(in *omsclient.OrderItemListReq) (*omsclient.OrderItemListResp, error) {
	all, err := l.svcCtx.OmsOrderItemModel.FindAll(l.ctx, in.Current, in.PageSize, 0)
	count, _ := l.svcCtx.OmsOrderItemModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询订单详情列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*omsclient.OrderItemListData
	for _, item := range *all {

		list = append(list, &omsclient.OrderItemListData{
			Id:                item.Id,
			OrderId:           item.OrderId,
			OrderSn:           item.OrderSn,
			ProductId:         item.ProductId,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductPrice:      float32(item.ProductPrice),
			ProductQuantity:   item.ProductQuantity,
			ProductSkuId:      item.ProductSkuId,
			ProductSkuCode:    item.ProductSkuCode,
			ProductCategoryId: item.ProductCategoryId,
			PromotionName:     item.PromotionName,
			PromotionAmount:   float32(item.PromotionAmount),
			CouponAmount:      float32(item.CouponAmount),
			IntegrationAmount: float32(item.IntegrationAmount),
			RealAmount:        float32(item.RealAmount),
			GiftIntegration:   item.GiftIntegration,
			GiftGrowth:        item.GiftGrowth,
			ProductAttr:       item.ProductAttr,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询订单详情列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &omsclient.OrderItemListResp{
		Total: count,
		List:  list,
	}, nil
}
