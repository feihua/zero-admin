package productservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductDetailByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductDetailByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductDetailByIdLogic {
	return &QueryProductDetailByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductDetailById 获取商品详情
// 1.获取商品信息
// 2.获取品牌信息
// 3.获取商品属性信息
// 4.获取商品属性值信息
// 5.获取商品SKU库存信息
// 6.商品阶梯价格设置
// 7.商品满减价格设置
// 8.获取商品的会员价格
func (l *QueryProductDetailByIdLogic) QueryProductDetailById(in *pmsclient.QueryProductDetailByIdReq) (*pmsclient.QueryProductDetailByIdResp, error) {
	// 1.获取商品信息
	productListData, product := buildProductListData(l, in.Id)

	productAttributeListData, attributeIds := buildProductAttributeListData(l, product)
	return &pmsclient.QueryProductDetailByIdResp{
		Product:                   productListData,
		Brand:                     buildBrandListData(l, product),
		ProductAttributeList:      productAttributeListData,
		ProductAttributeValueList: buildProductAttributeValueListData(l, product, attributeIds),
		SkuStockList:              buildSkuStockListData(l, product),
		ProductLadderList:         buildProductLadderListData(l, product),
		ProductFullReductionList:  buildProductFullReductionListData(l, product),
		MemberPriceList:           buildProductMemberListData(l, product),
	}, nil
}

// 1.获取商品信息
func buildProductListData(l *QueryProductDetailByIdLogic, productId int64) (*pmsclient.ProductListData, *model.PmsProduct) {
	item, _ := query.PmsProduct.WithContext(l.ctx).Where(query.PmsProduct.ID.Eq(productId)).First()

	return &pmsclient.ProductListData{
		Id:                         item.ID,                                      //
		BrandId:                    item.BrandID,                                 // 品牌id
		ProductCategoryId:          item.ProductCategoryID,                       // 商品分类id
		FeightTemplateId:           item.FeightTemplateID,                        // 商品运费模板id
		ProductAttributeCategoryId: item.ProductAttributeCategoryID,              // 商品属性分类id
		Name:                       item.Name,                                    // 商品名称
		Pic:                        item.Pic,                                     // 商品图片
		ProductSn:                  item.ProductSn,                               // 货号
		DeleteStatus:               item.DeleteStatus,                            // 删除状态：0->未删除；1->已删除
		PublishStatus:              item.PublishStatus,                           // 上架状态：0->下架；1->上架
		NewStatus:                  item.NewStatus,                               // 新品状态:0->不是新品；1->新品
		RecommandStatus:            item.RecommandStatus,                         // 推荐状态；0->不推荐；1->推荐
		VerifyStatus:               item.VerifyStatus,                            // 审核状态：0->未审核；1->审核通过
		Sort:                       item.Sort,                                    // 排序
		Sale:                       item.Sale,                                    // 销量
		Price:                      item.Price,                                   // 商品价格
		PromotionPrice:             item.PromotionPrice,                          // 促销价格
		GiftGrowth:                 item.GiftGrowth,                              // 赠送的成长值
		GiftPoint:                  item.GiftPoint,                               // 赠送的积分
		UsePointLimit:              item.UsePointLimit,                           // 限制使用的积分数
		SubTitle:                   item.SubTitle,                                // 副标题
		Description:                item.Description,                             // 商品描述
		OriginalPrice:              item.OriginalPrice,                           // 市场价
		Stock:                      item.Stock,                                   // 库存
		LowStock:                   item.LowStock,                                // 库存预警值
		Unit:                       item.Unit,                                    // 单位
		Weight:                     item.Weight,                                  // 商品重量，默认为克
		PreviewStatus:              item.PreviewStatus,                           // 是否为预告商品：0->不是；1->是
		ServiceIds:                 item.ServiceIds,                              // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
		Keywords:                   item.Keywords,                                // 搜索关键字
		Note:                       item.Note,                                    // 备注
		AlbumPics:                  item.AlbumPics,                               // 画册图片，连产品图片限制为5张，以逗号分割
		DetailTitle:                item.DetailTitle,                             // 详情标题
		DetailDesc:                 item.DetailDesc,                              // 详情描述
		DetailHtml:                 item.DetailHTML,                              // 产品详情网页内容
		DetailMobileHtml:           item.DetailMobileHTML,                        // 移动端网页详情
		PromotionStartTime:         time_util.TimeToStr(item.PromotionStartTime), // 促销开始时间
		PromotionEndTime:           time_util.TimeToStr(item.PromotionEndTime),   // 促销结束时间
		PromotionPerLimit:          item.PromotionPerLimit,                       // 活动限购数量
		PromotionType:              item.PromotionType,                           // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		BrandName:                  item.BrandName,                               // 品牌名称
		ProductCategoryName:        item.ProductCategoryName,                     // 商品分类名称
		ProductCategoryIdArray:     item.ProductCategoryIDArray,                  // 商品分类id字符串
	}, item
}

// 2.获取品牌信息
func buildBrandListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) *pmsclient.BrandData {
	item, _ := query.PmsBrand.WithContext(l.ctx).Where(query.PmsBrand.ID.Eq(pmsProduct.BrandID)).First()
	return &pmsclient.BrandData{
		Id:                  item.ID,
		Name:                item.Name,
		FirstLetter:         item.FirstLetter,
		Sort:                item.Sort,
		FactoryStatus:       item.FactoryStatus,
		ShowStatus:          item.ShowStatus,
		ProductCount:        item.ProductCount,
		ProductCommentCount: item.ProductCommentCount,
		Logo:                item.Logo,
		BigPic:              item.BigPic,
		BrandStory:          item.BrandStory,
	}
}

// 3.获取商品属性信息
func buildProductAttributeListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) ([]*pmsclient.ProductAttributeDataList, []int64) {

	q := query.PmsProductAttribute
	result, _ := q.WithContext(l.ctx).Where(q.Type.Eq(2), q.ProductAttributeCategoryID.Eq(pmsProduct.ProductAttributeCategoryID)).Find()

	var list []*pmsclient.ProductAttributeDataList
	var attributeIds []int64
	for _, item := range result {
		attributeIds = append(attributeIds, item.ID)
		list = append(list, &pmsclient.ProductAttributeDataList{
			Id:                         item.ID,                         //
			ProductAttributeCategoryId: item.ProductAttributeCategoryID, // 商品属性分类id
			Name:                       item.Name,                       // 商品属性分类id
			SelectType:                 item.SelectType,                 // 属性选择类型：0->唯一；1->单选；2->多选
			InputType:                  item.InputType,                  // 属性录入方式：0->手工录入；1->从列表中选取
			InputList:                  item.InputList,                  // 可选值列表，以逗号隔开
			Sort:                       item.Sort,                       // 排序字段：最高的可以单独上传图片
			FilterType:                 item.FilterType,                 // 分类筛选样式：1->普通；1->颜色
			SearchType:                 item.SearchType,                 // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
			RelatedStatus:              item.RelatedStatus,              // 相同属性产品是否关联；0->不关联；1->关联
			HandAddStatus:              item.HandAddStatus,              // 是否支持手动新增；0->不支持；1->支持
			Type:                       item.Type,                       // 属性的类型；0->规格；1->参数
		})
	}

	return list, attributeIds
}

// 4.获取商品属性值信息
func buildProductAttributeValueListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct, attributeIds []int64) []*pmsclient.ProductAttributeValueData {
	if len(attributeIds) > 0 {
		q := query.PmsProductAttributeValue
		result, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(pmsProduct.ID)).Find()

		var list []*pmsclient.ProductAttributeValueData
		for _, item := range result {
			list = append(list, &pmsclient.ProductAttributeValueData{
				Id:                 item.ID,                 //
				ProductId:          item.ProductID,          // 商品id
				ProductAttributeId: item.ProductAttributeID, // 商品属性id
				Value:              item.Value,              // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
			})
		}

		return list
	}

	return nil
}

// 5.获取商品SKU库存信息
func buildSkuStockListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) []*pmsclient.SkuStockData {
	q := query.PmsSkuStock
	result, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(pmsProduct.ID)).Find()
	var list []*pmsclient.SkuStockData
	for _, item := range result {

		list = append(list, &pmsclient.SkuStockData{
			Id:             item.ID,             //
			ProductId:      item.ProductID,      // 商品id
			SkuCode:        item.SkuCode,        // sku编码
			Price:          item.Price,          // 价格
			Stock:          item.Stock,          // 库存
			LowStock:       item.LowStock,       // 预警库存
			Pic:            item.Pic,            // 展示图片
			Sale:           item.Sale,           // 销量
			PromotionPrice: item.PromotionPrice, // 单品促销价格
			LockStock:      item.LockStock,      // 锁定库存
			SpData:         item.SpData,         // 商品销售属性，json格式
		})
	}

	return list
}

// 6.商品阶梯价格设置
func buildProductLadderListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) []*pmsclient.ProductLadderData {
	if pmsProduct.PromotionType == 3 {
		q := query.PmsProductLadder
		result, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(pmsProduct.ID)).Find()
		var list []*pmsclient.ProductLadderData
		for _, item := range result {

			list = append(list, &pmsclient.ProductLadderData{
				Id:        item.ID,        //
				ProductId: item.ProductID, // 商品id
				Count:     item.Count,     // 满足的商品数量
				Discount:  item.Discount,  // 折扣
				Price:     item.Price,     // 折后价格
			})
		}

		return list
	}

	return nil
}

// 7.商品满减价格设置
func buildProductFullReductionListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) []*pmsclient.ProductFullReductionData {
	if pmsProduct.PromotionType == 4 {
		q := query.PmsProductFullReduction
		result, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(pmsProduct.ID)).Find()

		var list []*pmsclient.ProductFullReductionData
		for _, item := range result {

			list = append(list, &pmsclient.ProductFullReductionData{
				Id:          item.ID,          //
				ProductId:   item.ProductID,   // 商品id
				FullPrice:   item.FullPrice,   // 商品满多少
				ReducePrice: item.ReducePrice, // 商品减多少
			})
		}
		return list
	}
	return nil

}

// 8.获取商品的会员价格
func buildProductMemberListData(l *QueryProductDetailByIdLogic, pmsProduct *model.PmsProduct) []*pmsclient.MemberPriceListData {
	if pmsProduct.PromotionType == 2 {
		q := query.PmsMemberPrice
		result, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(pmsProduct.ID)).Find()

		var list []*pmsclient.MemberPriceListData
		for _, item := range result {

			list = append(list, &pmsclient.MemberPriceListData{
				Id:              item.ID,              //
				ProductId:       item.ProductID,       // 商品id
				MemberLevelId:   item.MemberLevelID,   // 会员等级id
				MemberPrice:     item.MemberPrice,     // 会员价格
				MemberLevelName: item.MemberLevelName, // 会员等级名称
			})
		}
		return list
	}
	return nil

}
