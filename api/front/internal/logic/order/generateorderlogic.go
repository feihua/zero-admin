package order

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/front/internal/logic/cart"
	"github.com/feihua/zero-admin/api/front/internal/logic/member/coupon"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// GenerateOrderLogic
/*
Author: LiuFeiHua
Date: 2023/12/12 18:04
*/
type GenerateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateOrderLogic {
	return &GenerateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GenerateOrder 根据提交信息生成订单
// 1.获取购物车及优惠信息
// 2.生成下单商品信息
// 3.判断购物车中商品是否都有库存
// 4.判断是否使用了优惠券
// 5.判断是否使用积分
// 6.计算order_item的实付金额
// 7.进行库存锁定
// 8.根据商品合计、运费、活动优惠、优惠券、积分计算应付金额
// 9.转化为订单信息并插入数据库(删除购物车中的下单商品)
// 10.如果使用优惠券,更新优惠券使用状态
// 11.如果使用积分,需要扣除积分
// 12.发送延迟消息取消订单
func (l *GenerateOrderLogic) GenerateOrder(req *types.GenerateOrderReq) (*types.GenerateOrderResp, error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	memberInfo, _ := l.svcCtx.MemberService.QueryMemberById(l.ctx, &umsclient.MemberByIdReq{Id: memberId})
	//1.获取购物车及优惠信息
	cartPromotionItemList := cart.QueryCartListPromotion(req.CartIds, l.ctx, l.svcCtx)
	if len(cartPromotionItemList) == 0 {
		return result(1, "购物车还没有商品,请先添加商品到购物车!"), nil
	}
	//2.生成下单商品信息
	var flag = false //用于判断库存
	var totalAmount int64 = 0
	orderItemList := make([]*omsclient.OrderItemListData, 0)
	cartItemIds := make([]int64, 0) //方便提交订单成功,后删除购物车中的商品
	for _, item := range cartPromotionItemList {
		orderItemList = append(orderItemList, &omsclient.OrderItemListData{
			OrderId:           0,
			OrderSn:           "",
			ProductId:         item.ProductId,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductPic,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductPrice:      item.Price,
			ProductQuantity:   item.Quantity,
			ProductSkuId:      item.ProductSkuId,
			ProductSkuCode:    item.ProductSkuCode,
			ProductCategoryId: item.ProductCategoryId,
			PromotionName:     item.PromotionMessage,
			PromotionAmount:   item.ReduceAmount,
			CouponAmount:      0, //设置为0是表示默认不使用优惠券,后面会判断如果使用了优惠券,则会重新设置该值
			IntegrationAmount: 0, //设置为0是表示默认不使用积分,后面会判断如果使用了积分,则会重新设置该值
			RealAmount:        0,
			GiftIntegration:   item.Integration,
			GiftGrowth:        item.Growth,
			ProductAttr:       item.ProductAttr,
		})

		totalAmount = totalAmount + item.Price
		if item.RealStock <= 0 || item.RealStock < item.Quantity {
			flag = true
		}

		cartItemIds = append(cartItemIds, item.Id)
	}
	//3.判断购物车中商品是否都有库存
	if flag {
		return result(1, "库存不足，无法下单"), nil
	}
	//4.判断是否使用了优惠券
	//因为nt64传参数默认为0,所以判断为0的时候,表示没有使用优惠券
	if req.CouponId > 0 {
		var couponHistoryDetailListData *smsclient.CouponHistoryDetailListData
		enableList, _ := coupon.QueryCouponList(l.svcCtx, l.ctx, cartPromotionItemList)
		for _, item := range enableList {
			if item.CouponListData.Id == req.CouponId {
				couponHistoryDetailListData = item
			}
		}
		var couponData = couponHistoryDetailListData.CouponListData

		if couponData == nil {
			return result(1, "该优惠券不可用"), nil
		}
		useType := couponData.UseType
		if useType == 0 {
			//全场通用
			//(商品价格/可用商品总价)*优惠券面额
			for _, item := range orderItemList {
				item.CouponAmount = (item.ProductPrice / totalAmount) * couponData.Amount
			}
		} else if useType == 1 {
			//指定分类
			categoryIdList := make(map[int64]int64, 0)
			for _, item := range couponHistoryDetailListData.CategoryRelationList {
				categoryIdList[item.ProductCategoryId] = item.ProductCategoryId
			}
			orderItemListByCategoryId := make([]*omsclient.OrderItemListData, 0)
			var totalAmountByCategoryId int64
			for _, item := range orderItemList {
				_, ok := categoryIdList[item.ProductCategoryId]
				if ok {
					totalAmountByCategoryId = totalAmountByCategoryId + item.ProductPrice
					orderItemListByCategoryId = append(orderItemListByCategoryId, item)
				}
			}
			for _, item := range orderItemListByCategoryId {
				item.CouponAmount = (item.ProductPrice / totalAmountByCategoryId) * couponData.Amount
			}
		} else if useType == 2 {
			//指定商品
			productIdList := make(map[int64]int64, 0)
			for _, item := range couponHistoryDetailListData.ProductRelationList {
				productIdList[item.ProductId] = item.ProductId
			}
			orderItemListByProductId := make([]*omsclient.OrderItemListData, 0)
			var totalAmountByProductId int64
			for _, item := range orderItemList {
				_, ok := productIdList[item.ProductCategoryId]
				if ok {
					orderItemListByProductId = append(orderItemListByProductId, item)
				}
			}

			for _, item := range orderItemListByProductId {
				item.CouponAmount = (item.ProductPrice / totalAmountByProductId) * couponData.Amount
			}
		}

	}
	//5.判断是否使用积分
	if req.UseIntegration > 0 {
		//1.先计算出积分的总金额
		//1.1判断用户是否有这么多积分
		if req.UseIntegration > memberInfo.Integration {
			return result(1, "积分大于用户拥有有积分"), nil
		}
		//1.2根据积分使用规则判断是否可用
		//是否可与优惠券共用
		consumeSetting, _ := l.svcCtx.IntegrationConsumeSettingService.QueryIntegrationConsumeSettingById(l.ctx, &umsclient.QueryIntegrationConsumeSettingByIdReq{Id: 1})
		if req.CouponId > 0 && consumeSetting.CouponStatus == 0 {
			return result(1, "不可与优惠券共用"), nil
		}
		//1.3是否达到最低使用积分门槛
		if req.UseIntegration < consumeSetting.UseUnit {
			return result(1, "未达到最低使用积分门槛"), nil
		}
		//1.4是否超过订单抵用最高百分比
		integrationAmount := int64(req.UseIntegration / consumeSetting.UseUnit)
		var maxPercent = int64(consumeSetting.MaxPercentPerOrder / 100)
		if integrationAmount > maxPercent*totalAmount {
			return result(1, "超过订单抵用最高百分比"), nil
		}

		//2.可用情况下分摊到可用商品中
		for _, item := range orderItemList {
			item.IntegrationAmount = (item.ProductPrice / totalAmount) * integrationAmount
		}
	}
	//6.计算order_item的实付金额
	//原价-促销优惠-优惠券抵扣-积分抵扣
	for _, item := range orderItemList {
		item.RealAmount = item.ProductPrice - item.PromotionAmount - item.CouponAmount - item.IntegrationAmount
	}
	//7.进行库存锁定
	var data []*pmsclient.LockSkuStockLockData
	for _, item := range orderItemList {
		data = append(data, &pmsclient.LockSkuStockLockData{
			ProductSkuId:    item.ProductSkuId,
			ProductQuantity: item.ProductQuantity,
		})
	}
	_, err := l.svcCtx.SkuStockService.LockSkuStockLock(l.ctx, &pmsclient.LockSkuStockLockReq{
		Data: data,
	})
	if err != nil {
		return result(1, "锁定库存异常"), nil
	}
	//8.根据商品合计、运费、活动优惠、优惠券、积分计算应付金额
	//计算订单优惠(促销价、满减、阶梯价)
	var promotionAmount int64
	//获取订单促销信息
	var promotionInfo string
	//计算订单优惠券抵扣金额
	var couponAmount int64
	//计算积分抵扣金额
	var integrationAmount int64
	//计算该订单赠送的积分
	var giftIntegration int32
	//计算赠送成长值
	var giftGrowth int32
	for _, item := range orderItemList {
		promotionAmount = promotionAmount + item.PromotionAmount*int64(item.ProductQuantity)
		promotionInfo = promotionInfo + item.PromotionName
		couponAmount = couponAmount + item.CouponAmount*int64(item.ProductQuantity)
		integrationAmount = integrationAmount + item.IntegrationAmount*int64(item.ProductQuantity)
		giftIntegration = giftIntegration + item.GiftIntegration*item.ProductQuantity
		giftGrowth = giftGrowth + item.GiftGrowth*item.ProductQuantity
	}
	//计算订单应付金额
	payAmount := totalAmount - promotionAmount - couponAmount - integrationAmount
	address, err := l.svcCtx.MemberReceiveAddressService.MemberReceiveAddressQueryDetail(l.ctx, &umsclient.MemberReceiveAddressQueryDetailReq{
		UserId:    memberId,
		AddressID: req.MemberReceiveAddressId,
	})
	if err != nil {
		return result(1, "查询会员地址异常"), nil
	}
	//设置自动收货天数
	orderSettingList, err := l.svcCtx.OrderSettingService.OrderSettingList(l.ctx, &omsclient.OrderSettingListReq{
		Current:  1,
		PageSize: 10,
	})
	if err != nil {
		return result(1, "查询订单设置异常"), nil
	}
	orderInfo := &omsclient.OrderAddReq{
		CartItemIds:           cartItemIds,                              //购物车ids(创建订单成功后删除购物车中的商品)
		MemberId:              memberId,                                 //会员id
		CouponId:              req.CouponId,                             //优惠券id
		OrderSn:               "",                                       //生成订单号
		MemberUsername:        memberInfo.Username,                      //会员名称
		TotalAmount:           totalAmount,                              //订单总金额
		PayAmount:             payAmount,                                //订单应付金额
		FreightAmount:         0,                                        //运费金额
		PromotionAmount:       promotionAmount,                          //订单优惠
		IntegrationAmount:     integrationAmount,                        //积分抵扣金额
		CouponAmount:          couponAmount,                             //计算订单优惠券金额
		DiscountAmount:        0,                                        //管理员后台调整订单使用的折扣金额
		PayType:               0,                                        //支付方式：0->未支付；1->支付宝；2->微信
		SourceType:            1,                                        //订单来源：0->PC订单；1->app订单
		Status:                0,                                        //订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
		OrderType:             0,                                        //订单类型：0->正常订单；1->秒杀订单
		DeliveryCompany:       "",                                       //
		DeliverySn:            "",                                       //
		AutoConfirmDay:        orderSettingList.List[0].ConfirmOvertime, //自动收货天数
		Integration:           giftIntegration,                          //赠送积分
		Growth:                giftGrowth,                               //赠送成长值
		PromotionInfo:         promotionInfo,                            //获取订单促销信息
		BillType:              0,                                        //
		BillHeader:            "",                                       //
		BillContent:           "",                                       //
		BillReceiverPhone:     "",                                       //
		BillReceiverEmail:     "",                                       //
		ReceiverName:          address.MemberName,                       //收货人名称
		ReceiverPhone:         address.PhoneNumber,                      //收货人电话
		ReceiverPostCode:      address.PostCode,                         //邮政编码
		ReceiverProvince:      address.Province,                         //省份/直辖市
		ReceiverCity:          address.City,                             //城市
		ReceiverRegion:        address.Region,                           //区
		ReceiverDetailAddress: address.DetailAddress,                    //详细地址(街道)
		Note:                  "",                                       //
		ConfirmStatus:         0,                                        //确认收货状态：0->未确认；1->已确认
		DeleteStatus:          0,                                        //删除状态：0->未删除；1->已删除
		UseIntegration:        req.UseIntegration,                       //下单时使用的积分
		OrderItemList:         orderItemList,                            //
	}
	//9.转化为订单信息并插入数据库(插入order表和order_item表)
	orderAddResp, err := l.svcCtx.OrderService.OrderAdd(l.ctx, orderInfo)
	if err != nil {
		return result(1, "提交订单异常"), nil
	}
	//10.如果使用优惠券,更新优惠券使用状态
	//10.1修改sms_coupon_history表的use_status字段
	//10.2记得修改sms_coupon的use_count字段,下单的时候要加1,取消订单的时候,要减1
	if couponAmount > 0 {
		_, err = l.svcCtx.CouponHistoryService.UpdateCouponStatus(l.ctx, &smsclient.UpdateCouponStatusReq{
			MemberId:  memberId,
			UseStatus: 1,
			CouponId:  req.CouponId,
		})
		if err != nil {
			return result(1, "更新优惠券使用状态异常"), nil
		}
	}
	//11.如果使用积分,需要扣除积分
	i := memberInfo.Integration - req.UseIntegration
	_, err = l.svcCtx.MemberService.UpdateMemberIntegration(l.ctx, &umsclient.UpdateMemberIntegrationReq{Id: memberId, Integration: int64(i)})
	if err != nil {
		return result(1, "扣除积分异常"), nil
	}
	//12.发送延迟消息取消订单
	//todo 发送延迟消息取消订单
	return &types.GenerateOrderResp{
		Code:    0,
		Message: "下单成功",
		Data: types.GenerateOrderData{
			Id:                orderAddResp.Id,
			MemberId:          memberId,
			MemberUsername:    memberInfo.Username,
			TotalAmount:       orderInfo.TotalAmount,
			PayAmount:         orderInfo.PayAmount,
			FreightAmount:     orderInfo.FreightAmount,
			PromotionAmount:   orderInfo.PromotionAmount,
			IntegrationAmount: orderInfo.IntegrationAmount,
			CouponAmount:      orderInfo.CouponAmount,
			DiscountAmount:    orderInfo.DiscountAmount,
			PayType:           orderInfo.PayType,
			SourceType:        orderInfo.SourceType,
			Status:            orderInfo.Status,
			OrderType:         orderInfo.OrderType,
			Integration:       orderInfo.Integration,
			Growth:            orderInfo.Growth,
			PromotionInfo:     orderInfo.PromotionInfo,
			Note:              orderInfo.Note,
			UseIntegration:    orderInfo.UseIntegration,
		},
	}, nil
}

func result(code int64, message string) *types.GenerateOrderResp {
	return &types.GenerateOrderResp{
		Code:    code,
		Message: message,
	}
}
