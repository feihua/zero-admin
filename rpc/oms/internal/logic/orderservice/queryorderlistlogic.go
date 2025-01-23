package orderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderListLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 15:09
*/
type QueryOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderListLogic {
	return &QueryOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderList app端查询会员的订单列表信息
func (l *QueryOrderListLogic) QueryOrderList(in *omsclient.QueryOrderListReq) (*omsclient.OrderListResp, error) {
	q := query.OmsOrder.WithContext(l.ctx)

	if in.MemberId != 0 {
		q = q.Where(query.OmsOrder.MemberID.Eq(in.MemberId))
	}
	if in.Status != 6 {
		q = q.Where(query.OmsOrder.Status.Eq(in.Status))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询订单列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*omsclient.OrderListData
	for _, item := range result {

		list = append(list, &omsclient.OrderListData{
			Id:                    item.ID,                                   // 订单id
			MemberId:              item.MemberID,                             // 会员id
			CouponId:              item.CouponID,                             // 优惠券id
			OrderSn:               item.OrderSn,                              // 订单编号
			CreateTime:            time_util.TimeToStr(item.CreateTime),      // 提交时间
			MemberUsername:        item.MemberUsername,                       // 用户帐号
			TotalAmount:           item.TotalAmount,                          // 订单总金额
			PayAmount:             item.PayAmount,                            // 应付金额（实际支付金额）
			FreightAmount:         item.FreightAmount,                        // 运费金额
			PromotionAmount:       item.PromotionAmount,                      // 促销优化金额（促销价、满减、阶梯价）
			IntegrationAmount:     item.IntegrationAmount,                    // 积分抵扣金额
			CouponAmount:          item.CouponAmount,                         // 优惠券抵扣金额
			DiscountAmount:        item.DiscountAmount,                       // 管理员后台调整订单使用的折扣金额
			PayType:               item.PayType,                              // 支付方式：0->未支付；1->支付宝；2->微信
			SourceType:            item.SourceType,                           // 订单来源：0->PC订单；1->app订单
			Status:                item.Status,                               // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
			OrderType:             item.OrderType,                            // 订单类型：0->正常订单；1->秒杀订单
			DeliveryCompany:       item.DeliveryCompany,                      // 物流公司(配送方式)
			DeliverySn:            item.DeliverySn,                           // 物流单号
			AutoConfirmDay:        item.AutoConfirmDay,                       // 自动确认时间（天）
			Integration:           item.Integration,                          // 可以获得的积分
			Growth:                item.Growth,                               // 可以活动的成长值
			PromotionInfo:         item.PromotionInfo,                        // 活动信息
			BillType:              item.BillType,                             // 发票类型：0->不开发票；1->电子发票；2->纸质发票
			BillHeader:            item.BillHeader,                           // 发票抬头
			BillContent:           item.BillContent,                          // 发票内容
			BillReceiverPhone:     item.BillReceiverPhone,                    // 收票人电话
			BillReceiverEmail:     item.BillReceiverEmail,                    // 收票人邮箱
			ReceiverName:          item.ReceiverName,                         // 收货人姓名
			ReceiverPhone:         item.ReceiverPhone,                        // 收货人电话
			ReceiverPostCode:      item.ReceiverPostCode,                     // 收货人邮编
			ReceiverProvince:      item.ReceiverProvince,                     // 省份/直辖市
			ReceiverCity:          item.ReceiverCity,                         // 城市
			ReceiverRegion:        item.ReceiverRegion,                       // 区
			ReceiverDetailAddress: item.ReceiverDetailAddress,                // 详细地址
			Note:                  item.Note,                                 // 订单备注
			ConfirmStatus:         item.ConfirmStatus,                        // 确认收货状态：0->未确认；1->已确认
			DeleteStatus:          item.DeleteStatus,                         // 删除状态：0->未删除；1->已删除
			UseIntegration:        item.UseIntegration,                       // 下单时使用的积分
			PaymentTime:           time_util.TimeToString(item.PaymentTime),  // 支付时间
			DeliveryTime:          time_util.TimeToString(item.DeliveryTime), // 发货时间
			ReceiveTime:           time_util.TimeToString(item.ReceiveTime),  // 确认收货时间
			CommentTime:           time_util.TimeToString(item.CommentTime),  // 评价时间
			ModifyTime:            time_util.TimeToString(item.ModifyTime),   // 修改时间
			ItemListData:          queryItems(l, item.ID),
		})
	}

	logc.Infof(l.ctx, "查询订单列表信息,参数：%+v,响应：%+v", in, list)
	return &omsclient.OrderListResp{
		Total: 0,
		List:  list,
	}, nil

}

func queryItems(l *QueryOrderListLogic, OrderId int64) []*omsclient.OrderItemData {
	result, _ := query.OmsOrderItem.WithContext(l.ctx).Where(query.OmsOrderItem.OrderID.Eq(OrderId)).Find()
	var itemListData []*omsclient.OrderItemData
	for _, item := range result {
		itemListData = append(itemListData, &omsclient.OrderItemData{
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
	return itemListData
}
