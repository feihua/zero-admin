package cart

import (
	"context"
	"encoding/json"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/oms/omsclient"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CarItemtListPromotionLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 16:18
*/
type CarItemtListPromotionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCarItemtListPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarItemtListPromotionLogic {
	return &CarItemtListPromotionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CarItemtListPromotion 获取某个会员的购物车列表,包括促销信息
//1.获取会员购物车里面所有商品信息(根据会员id查询购物车所有商品,(表：oms_cart))
//2.判断是否会员选中购物车中的商品,如果没有就是会员购物车中所有有商品进行计算
//3.先根据productId对CartItem进行分组，以spu为单位进行计算优惠
//比如：购物车中有二台iphone15 pro max,其中一台是iphone15 pro max 是粉色,另外一台是黑色。{'productId':[iphone15 pro max 是粉色,iphone15 pro max 是黑色]}
//4.查询所有商品的优惠相关信息(商品的促销信息，包括sku、打折优惠、满减优惠)
//过滤出所有商品的id,查询表pms_product,pms_sku_stock,pms_product_ladder,pms_product_full_reduction
//5.根据商品促销类型计算商品促销优惠价格
//根据productId从第4步返回优惠列表中,查询出优惠信息
//6.根据表pms_product中的promotionType字段判断优惠类型:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
//6.1如果promotionType为0,没有促销使用原价
//6.2如果promotionType为1,单品促销(商品原价-促销价)
//从pms_sku_stock中获取price原价-单品促销价格pms_sku_stock中的promotionPrice字段,即可
//6.3如果promotionType为2,会员价格
//从pms_member_price表中查询
//6.4如果promotionType为3,打折优惠(商品原价-折扣*商品原价)
//从pms_sku_stock中获取price原价-折扣*商品原价(表pms_ladder),即可
//6.5如果promotionType为4,满减,
//6.5.1 计算分组中的总优惠金额
//6.5.2 (商品原价/总价)*满减金额(平摊到每件商品上的优惠金额)
// 6.6如果promotionType为5,限时购
//从sms_flash_promotion_product_relation表获取价格
func (l *CarItemtListPromotionLogic) CarItemtListPromotion(req *types.CarItemListPromotionReq) (resp *types.CarItemtListPromotionResp, err error) {
	//1.获取会员购物车里面所有商品信息(根据会员id查询购物车所有商品,(表：oms_cart))
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	cartItemListResp, _ := l.svcCtx.CartItemService.CartItemList(l.ctx, &omsclient.CartItemListReq{MemberId: memberId})

	//2.判断是否会员选中购物车中的商品,如果没有就是会员购物车中所有有商品进行计算
	cartItemListData := cartItemListResp.List
	itemListData := make([]*omsclient.CartItemListData, 0)
	if len(req.Ids) > 0 {
		for _, item := range cartItemListData {
			for _, id := range req.Ids {
				if item.Id == id {
					itemListData = append(itemListData, &omsclient.CartItemListData{
						Id:                item.Id,
						ProductId:         item.ProductId,
						ProductSkuId:      item.ProductSkuId,
						MemberId:          item.MemberId,
						Quantity:          item.Quantity,
						Price:             item.Price,
						ProductPic:        item.ProductPic,
						ProductName:       item.ProductName,
						ProductSubTitle:   item.ProductSubTitle,
						ProductSkuCode:    item.ProductSkuCode,
						MemberNickname:    item.MemberNickname,
						CreateDate:        item.CreateDate,
						ModifyDate:        item.ModifyDate,
						DeleteStatus:      item.DeleteStatus,
						ProductCategoryId: item.ProductCategoryId,
						ProductBrand:      item.ProductSn,
						ProductSn:         item.ProductSn,
						ProductAttr:       item.ProductAttr,
					})
				}
			}

		}

		//赋值用户选择购物车中的商品
		cartItemListData = itemListData
	}

	//3.先根据productId对CartItem进行分组，以spu为单位进行计算优惠(为后面满多少件打折做准备)
	//比如：购物车中有二台iphone15 pro max,其中一台是iphone15 pro max 是粉色,另外一台是黑色。{'productId':[iphone15 pro max 是粉色,iphone15 pro max 是黑色]}
	productCartMap := make(map[int64][]*omsclient.CartItemListData)
	for _, item := range cartItemListData {
		listData := productCartMap[item.ProductId]
		if listData == nil {
			listData = make([]*omsclient.CartItemListData, 0)
			listData = append(listData, item)
			productCartMap[item.ProductId] = listData
		} else {
			listData = append(listData, item)
		}
	}

	//4.查询所有商品的优惠相关信息(商品的促销信息，包括sku、打折优惠、满减优惠)
	//过滤出所有商品的id,查询表pms_product,pms_sku_stock,pms_product_ladder,pms_product_full_reduction
	productList := make([]*pmsclient.ProductDetailByIdResp, 0)            //方便后面统计金额
	productListMap := make(map[int64]*pmsclient.ProductDetailByIdResp, 0) //方便取数判断
	for _, item := range cartItemListData {
		productResp, _ := l.svcCtx.ProductService.ProductDetailById(l.ctx, &pmsclient.ProductDetailByIdReq{
			Id: item.ProductId,
		})
		productList = append(productList, productResp)
		productListMap[item.ProductId] = productResp
	}

	cartPromotionItemList := make([]types.CarItemtPromotionListData, 0)
	//5.根据商品促销类型计算商品促销优惠价格
	for productId, itemList := range productCartMap {
		//根据productId从第4步返回优惠列表中,查询出优惠信息
		promotionProduct := productListMap[productId]
		product := promotionProduct.Product
		skuStockList := promotionProduct.SkuStockList
		promotionType := product.PromotionType

		//6.根据表pms_product中的promotionType字段判断优惠类型:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购

		//6.1如果promotionType为0,没有促销使用原价
		if promotionType == 0 {
			cartPromotionItemList = handleNoReduce(itemList, skuStockList, product, cartPromotionItemList)
			//6.2如果promotionType为1,单品促销(商品原价-促销价)
			//从pms_sku_stock中获取price原价-单品促销价格pms_sku_stock中的promotionPrice字段,即可
		} else if promotionType == 1 {
			for _, item := range itemList {
				skuStock := getSkuStock(skuStockList, item.ProductSkuId)
				cartPromotionItem := types.CarItemtPromotionListData{
					Id:                item.Id,
					ProductId:         item.ProductId,
					ProductSkuId:      item.ProductSkuId,
					MemberId:          item.MemberId,
					Quantity:          item.Quantity,
					Price:             skuStock.Price, //单品促销使用原价
					ProductPic:        item.ProductPic,
					ProductName:       item.ProductName,
					ProductSubTitle:   item.ProductSubTitle,
					ProductSkuCode:    item.ProductSkuCode,
					MemberNickname:    item.MemberNickname,
					CreateDate:        item.CreateDate,
					ModifyDate:        item.ModifyDate,
					DeleteStatus:      item.DeleteStatus,
					ProductCategoryId: item.ProductCategoryId,
					ProductBrand:      item.ProductBrand,
					ProductSn:         item.ProductSn,
					ProductAttr:       item.ProductAttr,
					PromotionMessage:  "单品促销",
					ReduceAmount:      skuStock.Price - skuStock.PromotionPrice, //商品原价-促销价
					RealStock:         skuStock.Stock - skuStock.LockStock,
					Integration:       product.GiftPoint,
					Growth:            product.GiftGrowth,
				}
				cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
			}
			//6.3如果promotionType为2,会员价格
			//从pms_member_price表中查询
		} else if promotionType == 2 {
			//6.4如果promotionType为3,打折优惠(商品原价-折扣*商品原价)
			//从pms_sku_stock中获取price原价-折扣*商品原价(表pms_ladder),即可
		} else if promotionType == 3 {
			//6.5如果promotionType为4,满减,
			//6.5.1 计算分组中的总优惠金额
			//6.5.2 (商品原价/总价)*满减金额(平摊到每件商品上的优惠金额)
		} else if promotionType == 4 {
			//6.6如果promotionType为5,限时购
			//从sms_flash_promotion_product_relation表获取价格
		} else if promotionType == 4 {

		}

	}

	return
}

//对没满足优惠条件的商品进行处理
func handleNoReduce(itemList []*omsclient.CartItemListData, skuStockList []*pmsclient.SkuStockListData, product *pmsclient.ProductListData, cartPromotionItemList []types.CarItemtPromotionListData) []types.CarItemtPromotionListData {
	for _, item := range itemList {
		skuStock := getSkuStock(skuStockList, item.ProductSkuId)
		cartPromotionItem := types.CarItemtPromotionListData{
			Id:                item.Id,
			ProductId:         item.ProductId,
			ProductSkuId:      item.ProductSkuId,
			MemberId:          item.MemberId,
			Quantity:          item.Quantity,
			Price:             skuStock.Price,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductSubTitle:   item.ProductSubTitle,
			ProductSkuCode:    item.ProductSkuCode,
			MemberNickname:    item.MemberNickname,
			CreateDate:        item.CreateDate,
			ModifyDate:        item.ModifyDate,
			DeleteStatus:      item.DeleteStatus,
			ProductCategoryId: item.ProductCategoryId,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductAttr:       item.ProductAttr,
			PromotionMessage:  "无优惠",
			ReduceAmount:      0,
			RealStock:         skuStock.Stock - skuStock.LockStock,
			Integration:       product.GiftPoint,
			Growth:            product.GiftGrowth,
		}
		cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
	}
	return cartPromotionItemList
}

////获取sku
func getSkuStock(skuStockList []*pmsclient.SkuStockListData, productSkuId int64) *pmsclient.SkuStockListData {
	var skuStock *pmsclient.SkuStockListData
	for _, sku := range skuStockList {
		if sku.Id == productSkuId {
			skuStock = sku
		}
	}
	return skuStock
}
