package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

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

func (l *OrderItemListLogic) OrderItemList(in *oms.OrderItemListReq) (*oms.OrderItemListResp, error) {
	all, err := l.svcCtx.OmsOrderItemModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.OmsOrderItemModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询订单详情列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

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

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询订单详情列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &oms.OrderItemListResp{
		Total: count,
		List:  list,
	}, nil
}
