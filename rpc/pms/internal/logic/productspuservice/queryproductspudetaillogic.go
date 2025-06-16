package productspuservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryProductSpuDetailLogic 查询商品SPU详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:38
*/
type QueryProductSpuDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductSpuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSpuDetailLogic {
	return &QueryProductSpuDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductSpuDetail 查询商品SPU详情
// 1.获取商品信息
// 2.获取品牌信息
// 3.获取商品属性信息
// 4.获取商品属性值信息
// 5.获取商品SKU库存信息
// 6.商品阶梯价格设置
// 7.商品满减价格设置
// 8.获取商品的会员价格
func (l *QueryProductSpuDetailLogic) QueryProductSpuDetail(in *pmsclient.QueryProductSpuDetailReq) (*pmsclient.QueryProductSpuDetailResp, error) {
	item, err := query.PmsProductSpu.WithContext(l.ctx).Where(query.PmsProductSpu.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品SPU不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品SPU不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品SPU异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品SPU异常")
	}

	product := &pmsclient.ProductSpuListData{
		Id:                  item.ID,                                          // 商品SpuId
		Name:                item.Name,                                        // 商品名称
		CategoryId:          item.CategoryID,                                  // 商品分类ID
		CategoryIds:         item.CategoryIds,                                 // 商品分类ID集合
		CategoryName:        item.CategoryName,                                // 商品分类名称
		BrandId:             item.BrandID,                                     // 品牌ID
		BrandName:           item.BrandName,                                   // 品牌名称
		Unit:                item.Unit,                                        // 单位
		Weight:              float32(item.Weight),                             // 重量(kg)
		Keywords:            item.Keywords,                                    // 关键词
		Brief:               item.Brief,                                       // 简介
		Description:         item.Description,                                 // 详细描述
		AlbumPics:           item.AlbumPics,                                   // 画册图片，最多8张，以逗号分割
		MainPic:             item.MainPic,                                     // 主图
		PriceRange:          item.PriceRange,                                  // 价格区间
		PublishStatus:       item.PublishStatus,                               // 上架状态：0-下架，1-上架
		NewStatus:           item.NewStatus,                                   // 新品状态:0->不是新品；1->新品
		RecommendStatus:     item.RecommendStatus,                             // 推荐状态；0->不推荐；1->推荐
		VerifyStatus:        item.VerifyStatus,                                // 审核状态：0->未审核；1->审核通过
		PreviewStatus:       item.PreviewStatus,                               // 是否为预告商品：0->不是；1->是
		Sort:                item.Sort,                                        // 排序
		NewStatusSort:       item.NewStatusSort,                               // 新品排序
		RecommendStatusSort: item.RecommendStatusSort,                         // 推荐排序
		Sales:               item.Sales,                                       // 销量
		Stock:               item.Stock,                                       // 库存
		LowStock:            item.LowStock,                                    // 预警库存
		PromotionType:       item.PromotionType,                               // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀
		DetailTitle:         item.DetailTitle,                                 // 详情标题
		DetailDesc:          item.DetailDesc,                                  // 详情描述
		DetailHtml:          item.DetailHTML,                                  // 产品详情网页内容
		DetailMobileHtml:    item.DetailMobileHTML,                            // 移动端网页详情
		CreateBy:            item.CreateBy,                                    // 创建人ID
		CreateTime:          time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:            pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:          time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	productAttributeListData, attributeIds := buildProductAttributeListData(l, item)
	return &pmsclient.QueryProductSpuDetailResp{
		Data:                      product,
		Brand:                     buildBrandListData(l, item),
		ProductAttributeList:      productAttributeListData,
		ProductAttributeValueList: buildProductAttributeValueListData(l, item, attributeIds),
		SkuStockList:              buildSkuStockListData(l, item),
		ProductLadderList:         buildProductLadderListData(l, item),
		ProductFullReductionList:  buildProductFullReductionListData(l, item),
		MemberPriceList:           buildProductMemberListData(l, item),
	}, nil
}

// 2.获取品牌信息
func buildBrandListData(l *QueryProductSpuDetailLogic, pmsProduct *model.PmsProductSpu) *pmsclient.BrandData {
	item, _ := query.PmsProductBrand.WithContext(l.ctx).Where(query.PmsProductBrand.ID.Eq(pmsProduct.BrandID)).First()
	return &pmsclient.BrandData{
		Id:                  item.ID,                                          //
		Name:                item.Name,                                        // 品牌名称
		Logo:                item.Logo,                                        // 品牌logo
		BigPic:              item.BigPic,                                      // 专区大图
		Description:         item.Description,                                 // 描述
		FirstLetter:         item.FirstLetter,                                 // 首字母
		Sort:                item.Sort,                                        // 排序
		RecommendStatus:     item.RecommendStatus,                             // 推荐状态
		ProductCount:        item.ProductCount,                                // 产品数量
		ProductCommentCount: item.ProductCommentCount,                         // 产品评论数量
		IsEnabled:           item.IsEnabled,                                   // 是否启用
		CreateBy:            item.CreateBy,                                    // 创建人ID
		CreateTime:          time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:            pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:          time_util.TimeToString(item.UpdateTime),          // 更新时间
	}
}

// 3.获取商品属性信息
func buildProductAttributeListData(l *QueryProductSpuDetailLogic, pmsProduct *model.PmsProductSpu) ([]*pmsclient.ProductAttributeDataList, []int64) {

	group, _ := query.PmsProductAttributeGroup.WithContext(l.ctx).Where(query.PmsProductAttributeGroup.CategoryID.Eq(pmsProduct.CategoryID)).First()
	q := query.PmsProductAttribute
	result, _ := q.WithContext(l.ctx).Where(q.GroupID.Eq(group.CategoryID)).Find()

	var list []*pmsclient.ProductAttributeDataList
	var attributeIds []int64
	for _, item := range result {
		attributeIds = append(attributeIds, item.ID)
		list = append(list, &pmsclient.ProductAttributeDataList{
			Id:           item.ID,                                          // 主键id
			GroupId:      item.GroupID,                                     // 属性分组ID
			Name:         item.Name,                                        // 属性名称
			InputType:    item.InputType,                                   // 输入类型：1-手动输入，2-单选，3-多选
			ValueType:    item.ValueType,                                   // 值类型：1-文本，2-数字，3-日期
			InputList:    item.InputList,                                   // 可选值列表，用逗号分隔
			Unit:         item.Unit,                                        // 单位
			IsRequired:   item.IsRequired,                                  // 是否必填
			IsSearchable: item.IsSearchable,                                // 是否支持搜索
			IsShow:       item.IsShow,                                      // 是否显示
			Sort:         item.Sort,                                        // 排序
			CreateBy:     item.CreateBy,                                    // 创建人ID
			CreateTime:   time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:     pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:   time_util.TimeToString(item.UpdateTime),          // 更新时间
		})
	}

	return list, attributeIds
}

// 4.获取商品属性值信息
func buildProductAttributeValueListData(l *QueryProductSpuDetailLogic, pmsProduct *model.PmsProductSpu, attributeIds []int64) []*pmsclient.ProductAttributeValueData {
	if len(attributeIds) > 0 {
		q := query.PmsProductAttributeValue
		result, _ := q.WithContext(l.ctx).Where(q.SpuID.Eq(pmsProduct.ID)).Find()

		var list []*pmsclient.ProductAttributeValueData
		for _, item := range result {
			list = append(list, &pmsclient.ProductAttributeValueData{
				Id:          item.ID,                                          // 主键id
				SpuId:       item.SpuID,                                       // 商品SPU ID
				AttributeId: item.AttributeID,                                 // 属性ID
				Value:       item.Value,                                       // 属性值
				CreateBy:    item.CreateBy,                                    // 创建人ID
				CreateTime:  time_util.TimeToStr(item.CreateTime),             // 创建时间
				UpdateBy:    pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
				UpdateTime:  time_util.TimeToString(item.UpdateTime),          // 更新时间
			})
		}

		return list
	}

	return nil
}

// 5.获取商品SKU库存信息
func buildSkuStockListData(l *QueryProductSpuDetailLogic, pmsProduct *model.PmsProductSpu) []*pmsclient.SkuStockData {
	q := query.PmsProductSku
	result, _ := q.WithContext(l.ctx).Where(q.SpuID.Eq(pmsProduct.ID)).Find()
	var list []*pmsclient.SkuStockData
	for _, item := range result {

		list = append(list, &pmsclient.SkuStockData{
			Id:                 item.ID,                                          // 商品SpuId
			SpuId:              item.SpuID,                                       // 商品SpuId
			Name:               item.Name,                                        // SKU名称
			SkuCode:            item.SkuCode,                                     // SKU编码
			MainPic:            item.MainPic,                                     // 主图
			AlbumPics:          item.AlbumPics,                                   // 图片集
			Price:              float32(item.Price),                              // 价格
			PromotionPrice:     float32(item.PromotionPrice),                     // 单品促销价格
			PromotionStartTime: time_util.TimeToString(item.PromotionStartTime),  // 促销开始时间
			PromotionEndTime:   time_util.TimeToString(item.PromotionEndTime),    // 促销结束时间
			Stock:              item.Stock,                                       // 库存
			LowStock:           item.LowStock,                                    // 预警库存
			SpecData:           item.SpecData,                                    // 规格数据
			Weight:             float32(item.Weight),                             // 重量(kg)
			PublishStatus:      item.PublishStatus,                               // 上架状态：0-下架，1-上架
			VerifyStatus:       item.VerifyStatus,                                // 审核状态：0-未审核，1-审核通过，2-审核不通过
			Sort:               item.Sort,                                        // 排序
			Sales:              item.Sales,                                       // 销量
			CreateBy:           item.CreateBy,                                    // 创建人ID
			CreateTime:         time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:           pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:         time_util.TimeToString(item.UpdateTime),          // 更新时间
		})
	}

	return list
}

// 6.商品阶梯价格设置
func buildProductLadderListData(l *QueryProductSpuDetailLogic, pmsProduct *model.PmsProductSpu) []*pmsclient.ProductLadderData {
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
func buildProductFullReductionListData(l *QueryProductSpuDetailLogic, pmsProduct *model.PmsProductSpu) []*pmsclient.ProductFullReductionData {
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
func buildProductMemberListData(l *QueryProductSpuDetailLogic, pmsProduct *model.PmsProductSpu) []*pmsclient.MemberPriceListData {
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
