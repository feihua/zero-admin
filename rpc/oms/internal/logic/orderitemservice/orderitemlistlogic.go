package orderitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	q := query.OmsOrderItem.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询订单详情列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*omsclient.OrderItemListData
	for _, item := range result {

		list = append(list, &omsclient.OrderItemListData{
			Id:                item.ID,
			OrderId:           item.OrderID,
			OrderSn:           item.OrderSn,
			ProductId:         item.ProductID,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductPrice:      item.ProductPrice,
			ProductQuantity:   item.ProductQuantity,
			ProductSkuId:      item.ProductSkuID,
			ProductSkuCode:    item.ProductSkuCode,
			ProductCategoryId: item.ProductCategoryID,
			PromotionName:     item.PromotionName,
			PromotionAmount:   item.PromotionAmount,
			CouponAmount:      item.CouponAmount,
			IntegrationAmount: item.IntegrationAmount,
			RealAmount:        item.RealAmount,
			GiftIntegration:   item.GiftIntegration,
			GiftGrowth:        item.GiftGrowth,
			ProductAttr:       item.ProductAttr,
		})
	}

	logc.Infof(l.ctx, "查询订单详情列表信息,参数：%+v,响应：%+v", in, list)
	return &omsclient.OrderItemListResp{
		Total: count,
		List:  list,
	}, nil
}
