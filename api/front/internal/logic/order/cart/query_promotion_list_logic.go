package cart

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strconv"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPromotionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPromotionListLogic {
	return &QueryPromotionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryPromotionList 获取某个会员的购物车列表,包括促销信息
// 1.获取会员购物车里面所有商品信息(根据会员id查询购物车所有商品,(表：oms_cart))
// 2.判断是否会员选中购物车中的商品,如果没有就是会员购物车中所有有商品进行计算
// 3.先根据productId对CartItem进行分组，以spu为单位进行计算优惠
// 比如：购物车中有二台iphone15 pro max,其中一台是iphone15 pro max 是粉色,另外一台是黑色。{'productId':[iphone15 pro max 是粉色,iphone15 pro max 是黑色]}
// 4.查询所有商品的优惠相关信息(商品的促销信息，包括sku、打折优惠、满减优惠)
// 过滤出所有商品的id,查询表pms_product,pms_sku_stock,pms_product_ladder,pms_product_full_reduction
// 5.根据商品促销类型计算商品促销优惠价格
// 根据productId从第4步返回优惠列表中,查询出优惠信息
// 6.根据表pms_product中的promotionType字段判断优惠类型:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
// 6.1如果promotionType为0,没有促销使用原价
// 6.2如果promotionType为1,单品促销(商品原价-促销价)
// 从pms_sku_stock中获取price原价-单品促销价格pms_sku_stock中的promotionPrice字段,即可
// 6.3如果promotionType为2,会员价格
// 从pms_member_price表中查询
// 6.4如果promotionType为3,打折优惠(商品原价-折扣*商品原价)
// 从pms_sku_stock中获取price原价-折扣*商品原价(表pms_ladder),即可
// 6.5如果promotionType为4,满减,
// 6.5.1 计算分组中的总优惠金额
// 6.5.2 (商品原价/总价)*满减金额(平摊到每件商品上的优惠金额)
// 6.6如果promotionType为5,限时购
// 从sms_flash_promotion_product_relation表获取价格
func (l *QueryPromotionListLogic) QueryPromotionList(req *types.CarItemListPromotionReq) (resp *types.CarItemtListPromotionResp, err error) {
	cartPromotionItemList, err := QueryCartListPromotion(req.Ids, l.ctx, l.svcCtx)

	if err != nil {
		return nil, err
	}

	return &types.CarItemtListPromotionResp{
		Code:    0,
		Message: "操作成功",
		Data:    cartPromotionItemList,
	}, nil
}

// QueryCartListPromotion 抽取单独方法(方便在下单确认的使用)
func QueryCartListPromotion(ids []int64, ctx context.Context, svcCtx *svc.ServiceContext) ([]types.CarItemtPromotionListData, error) {
	// 1.获取会员购物车里面所有商品信息(根据会员id查询购物车所有商品,(表：oms_cart))
	memberId, _ := ctx.Value("memberId").(json.Number).Int64()
	cartItemListResp, err := svcCtx.CartItemService.QueryCartItemList(ctx, &omsclient.QueryCartItemListReq{MemberId: memberId})

	if err != nil {
		logc.Errorf(ctx, "获取会员购物车里面所有商品信息失败,参数memberId: %+v,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 2.判断是否会员选中购物车中的商品,如果没有就是会员购物车中所有有商品进行计算
	cartItemListData := cartItemListResp.List
	itemListData := make([]*omsclient.CartItemData, 0)
	if len(ids) > 0 {
		for _, item := range cartItemListData {
			for _, id := range ids {
				if item.Id == id {
					cartItemListData1 := omsclient.CartItemData{}
					_ = copier.Copy(&cartItemListData1, &item)
					itemListData = append(itemListData, &cartItemListData1)
				}
			}
		}
		// 赋值用户选择购物车中的商品
		cartItemListData = itemListData
	}

	// 3.先根据productId对CartItem进行分组，以spu为单位进行计算优惠(为后面满多少件打折做准备)
	// 比如：购物车中有二台iphone15 pro max,其中一台是iphone15 pro max 是粉色,另外一台是黑色。{'productId':[iphone15 pro max 是粉色,iphone15 pro max 是黑色]}
	productCartMap := make(map[int64][]*omsclient.CartItemData)
	for _, item := range cartItemListData {
		listData := productCartMap[item.ProductId]
		if listData == nil {
			listData = make([]*omsclient.CartItemData, 0)
			listData = append(listData, item)
			productCartMap[item.ProductId] = listData
		} else {
			listData = append(listData, item)
		}
	}

	// 4.查询所有商品的优惠相关信息(商品的促销信息，包括sku、打折优惠、满减优惠)
	// 过滤出所有商品的id,查询表pms_product,pms_sku_stock,pms_product_ladder,pms_product_full_reduction
	productList := make([]*pmsclient.QueryProductSpuDetailResp, 0)            // 方便后面统计金额
	productListMap := make(map[int64]*pmsclient.QueryProductSpuDetailResp, 0) // 方便取数判断
	for _, item := range cartItemListData {
		productResp, _ := svcCtx.ProductSpuService.QueryProductSpuDetail(ctx, &pmsclient.QueryProductSpuDetailReq{
			Id: item.ProductId,
		})
		productList = append(productList, productResp)
		productListMap[item.ProductId] = productResp
	}

	cartPromotionItemList := make([]types.CarItemtPromotionListData, 0)
	// 5.根据商品促销类型计算商品促销优惠价格
	for productId, itemList := range productCartMap {
		// 根据productId从第4步返回优惠列表中,查询出优惠信息
		promotionProduct := productListMap[productId]
		product := promotionProduct.Data
		skuStockList := promotionProduct.SkuStockList
		promotionType := product.PromotionType

		// 6.根据表pms_product中的promotionType字段判断优惠类型:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购

		// 6.1如果promotionType为0,没有促销使用原价
		if promotionType == 0 {
			cartPromotionItemList = handleNoReduce(itemList, skuStockList, product, cartPromotionItemList)
			// 6.2如果promotionType为1,单品促销(商品原价-促销价)
			// 从pms_sku_stock中获取price原价-单品促销价格pms_sku_stock中的promotionPrice字段,即可
		} else if promotionType == 1 {
			for _, item := range itemList {
				skuStock := getSkuStock(skuStockList, item.ProductSkuId)
				cartPromotionItem := types.CarItemtPromotionListData{}
				_ = copier.Copy(&cartPromotionItem, &item)
				cartPromotionItem.Price = skuStock.Price // 单品促销使用原价
				cartPromotionItem.PromotionMessage = "单品促销"
				cartPromotionItem.ReduceAmount = int64(skuStock.Price - skuStock.PromotionPrice) // 商品原价-促销价
				// cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
				// cartPromotionItem.Integration = product.GiftPoint
				// cartPromotionItem.Growth = product.GiftGrowth
				cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
			}
			// 6.3如果promotionType为2,会员价格
			// 从pms_member_price表中查询
		} else if promotionType == 2 {
			// 获取会员等级
			member, _ := svcCtx.MemberService.QueryMemberInfoDetail(ctx, &umsclient.QueryMemberInfoDetailReq{
				MemberId: memberId,
			})
			// 获取会员价格
			var memberPrice int64
			memberPriceList := promotionProduct.MemberPriceList
			for _, item := range memberPriceList {
				if item.MemberLevelId == member.LevelId {
					memberPrice = item.MemberPrice
				}
			}
			for _, item := range itemList {
				skuStock := getSkuStock(skuStockList, item.ProductSkuId)
				cartPromotionItem := types.CarItemtPromotionListData{}
				_ = copier.Copy(&cartPromotionItem, &item)
				cartPromotionItem.Price = skuStock.Price // 单品促销使用原价
				cartPromotionItem.PromotionMessage = "会员价格"
				cartPromotionItem.ReduceAmount = memberPrice // 会员价
				// cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
				// cartPromotionItem.Integration = product.GiftPoint
				// cartPromotionItem.Growth = product.GiftGrowth
				cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
			}
			// 6.4如果promotionType为3,打折优惠(商品原价-折扣*商品原价)
			// 从pms_sku_stock中获取price原价-折扣*商品原价(表pms_ladder),即可
		} else if promotionType == 3 {
			// 获取购物车中指定商品的数量
			var count int32 = 0
			for _, item := range itemList {
				count = count + item.Quantity
			}

			// 根据购买商品数量获取满足条件的打折优惠策略
			productLadderList := promotionProduct.ProductLadderList
			var productLadder *pmsclient.ProductLadderData
			for _, item := range productLadderList {
				if count > item.Count {
					productLadder = item
				}
			}
			if productLadder != nil {
				for _, item := range itemList {
					skuStock := getSkuStock(skuStockList, item.ProductSkuId)
					cartPromotionItem := types.CarItemtPromotionListData{}
					_ = copier.Copy(&cartPromotionItem, &item)
					cartPromotionItem.Price = skuStock.Price
					c := strconv.FormatInt(int64(productLadder.Count), 10)
					d := fmt.Sprintf("%1.0f", productLadder.Discount*10)
					cartPromotionItem.PromotionMessage = "打折优惠：满" + c + "件,打" + d + "折"
					// 商品原价-折扣*商品原价
					// cartPromotionItem.ReduceAmount = skuStock.Price - productLadder.Discount*skuStock.Price
					// cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
					// cartPromotionItem.Integration = product.GiftPoint
					// cartPromotionItem.Growth = product.GiftGrowth
					cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
				}
			} else {
				cartPromotionItemList = handleNoReduce(itemList, skuStockList, product, cartPromotionItemList)
			}
			// 6.5如果promotionType为4,满减,
			// 6.5.1 计算分组中的总优惠金额
			// 6.5.2 (商品原价/总价)*满减金额(平摊到每件商品上的优惠金额)
		} else if promotionType == 4 {
			// 获取购物车中指定商品的总价
			var amount int64 = 0
			// for _, item := range itemList {
			// 	skuStock := getSkuStock(skuStockList, item.ProductSkuId)
			// 	amount = amount + skuStock.Price*int64(item.Quantity)
			//
			// }
			// 获取满减
			productFullReductionList := promotionProduct.ProductFullReductionList
			var productFull *pmsclient.ProductFullReductionData
			for _, item := range productFullReductionList {
				if amount-item.FullPrice > 0 {
					productFull = item
				}
			}
			if productFull != nil {
				for _, item := range itemList {
					skuStock := getSkuStock(skuStockList, item.ProductSkuId)
					cartPromotionItem := types.CarItemtPromotionListData{}
					_ = copier.Copy(&cartPromotionItem, &item)
					cartPromotionItem.Price = skuStock.Price
					f := fmt.Sprintf("%d", productFull.FullPrice)
					r := fmt.Sprintf("%d", productFull.ReducePrice)
					cartPromotionItem.PromotionMessage = "满减优惠：满" + f + "元,减" + r + "元"
					// (商品原价/总价)*满减金额
					// cartPromotionItem.ReduceAmount = (skuStock.Price / amount) * productFull.ReducePrice
					// cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
					// cartPromotionItem.Integration = product.GiftPoint
					// cartPromotionItem.Growth = product.GiftGrowth
					cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
				}
			} else {
				cartPromotionItemList = handleNoReduce(itemList, skuStockList, product, cartPromotionItemList)
			}
			// 6.6如果promotionType为5,限时购
			// 从sms_flash_promotion_product_relation表获取价格
		} else if promotionType == 5 {
			for _, item := range itemList {
				promotionByProduct, _ := svcCtx.SeckillProductService.QuerySeckillProductBySkuId(ctx, &smsclient.QuerySeckillProductBySkuIdReq{
					SkuId: item.ProductId,
				})
				skuStock := getSkuStock(skuStockList, item.ProductSkuId)
				cartPromotionItem := types.CarItemtPromotionListData{}
				_ = copier.Copy(&cartPromotionItem, &item)
				cartPromotionItem.Price = skuStock.Price
				cartPromotionItem.PromotionMessage = "限时购"
				cartPromotionItem.ReduceAmount = int64(promotionByProduct.SeckillPrice)
				// cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
				// cartPromotionItem.Integration = product.GiftPoint
				// cartPromotionItem.Growth = product.GiftGrowth
				cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
			}
		}

	}
	return cartPromotionItemList, nil
}

// 对没满足优惠条件的商品进行处理
func handleNoReduce(itemList []*omsclient.CartItemData, skuStockList []*pmsclient.SkuStockData, product *pmsclient.ProductSpuListData, cartPromotionItemList []types.CarItemtPromotionListData) []types.CarItemtPromotionListData {
	for _, item := range itemList {
		skuStock := getSkuStock(skuStockList, item.ProductSkuId)
		cartPromotionItem := types.CarItemtPromotionListData{}
		_ = copier.Copy(&cartPromotionItem, &item)
		cartPromotionItem.Price = skuStock.Price
		cartPromotionItem.PromotionMessage = "无优惠"
		cartPromotionItem.ReduceAmount = 0
		// cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
		// cartPromotionItem.Integration = product.GiftPoint
		// cartPromotionItem.Growth = product.GiftGrowth
		cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
	}
	return cartPromotionItemList
}

// 获取sku
func getSkuStock(skuStockList []*pmsclient.SkuStockData, productSkuId int64) *pmsclient.SkuStockData {
	var skuStock *pmsclient.SkuStockData
	for _, sku := range skuStockList {
		if sku.Id == productSkuId {
			skuStock = sku
		}
	}
	return skuStock
}
