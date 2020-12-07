package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/oms/internal/svc"
	"go-zero-admin/rpc/oms/oms"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *OrderItemListLogic) OrderItemList(in *oms.OrderItemListReq) (*oms.OrderItemListResp, error) {
	all, _ := l.svcCtx.OmsOrderItemModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*oms.OrderItemListData
	for _, item := range *all {

		list = append(list, &oms.OrderItemListData{
			Id:                item.Id,
			OrderId:           item.OrderId,
			OrderSn:           item.OrderSn,
			ProductId:         item.ProductId,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductPrice:      int64(item.ProductPrice),
			ProductQuantity:   item.ProductQuantity,
			ProductSkuId:      item.ProductSkuId,
			ProductSkuCode:    item.ProductSkuCode,
			ProductCategoryId: item.ProductCategoryId,
			PromotionName:     item.PromotionName,
			PromotionAmount:   int64(item.PromotionAmount),
			CouponAmount:      int64(item.CouponAmount),
			IntegrationAmount: int64(item.IntegrationAmount),
			RealAmount:        int64(item.RealAmount),
			GiftIntegration:   item.GiftIntegration,
			GiftGrowth:        item.GiftGrowth,
			ProductAttr:       item.ProductAttr,
		})
	}

	fmt.Println(list)
	return &oms.OrderItemListResp{
		Total: 10,
		List:  list,
	}, nil
}
