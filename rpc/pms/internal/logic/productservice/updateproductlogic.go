package productservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"math/rand"
	"strconv"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductLogic 更新商品
/*
Author: LiuFeiHua
Date: 2024/6/12 17:16
*/
type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProduct 更新商品
// 逻辑步骤：
// 1.更新商品基本信息
// 2.会员价格
// 3.阶梯价格
// 4.满减价格
// 5.更新sku库存信息
// 6.更新商品参数,添加自定义商品规格
func (l *UpdateProductLogic) UpdateProduct(in *pmsclient.UpdateProductReq) (*pmsclient.UpdateProductResp, error) {
	PromotionStartTime, _ := time.Parse("2006-01-02 15:04:05", in.PromotionStartTime)
	PromotionEndTime, _ := time.Parse("2006-01-02 15:04:05", in.PromotionEndTime)

	productId := in.Id

	q := query.PmsProduct
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProduct{
		ID:                         in.Id,                         //
		BrandID:                    in.BrandId,                    // 品牌id
		ProductCategoryID:          in.ProductCategoryId,          // 商品分类id
		FeightTemplateID:           in.FeightTemplateId,           // 商品运费模板id
		ProductAttributeCategoryID: in.ProductAttributeCategoryId, // 商品属性分类id
		Name:                       in.Name,                       // 商品名称
		Pic:                        in.Pic,                        // 商品图片
		ProductSn:                  in.ProductSn,                  // 货号
		DeleteStatus:               in.DeleteStatus,               // 删除状态：0->未删除；1->已删除
		PublishStatus:              in.PublishStatus,              // 上架状态：0->下架；1->上架
		NewStatus:                  in.NewStatus,                  // 新品状态:0->不是新品；1->新品
		RecommandStatus:            in.RecommandStatus,            // 推荐状态；0->不推荐；1->推荐
		VerifyStatus:               in.VerifyStatus,               // 审核状态：0->未审核；1->审核通过
		Sort:                       in.Sort,                       // 排序
		Sale:                       in.Sale,                       // 销量
		Price:                      in.Price,                      // 商品价格
		PromotionPrice:             in.PromotionPrice,             // 促销价格
		GiftGrowth:                 in.GiftGrowth,                 // 赠送的成长值
		GiftPoint:                  in.GiftPoint,                  // 赠送的积分
		UsePointLimit:              in.UsePointLimit,              // 限制使用的积分数
		SubTitle:                   in.SubTitle,                   // 副标题
		Description:                in.Description,                // 商品描述
		OriginalPrice:              in.OriginalPrice,              // 市场价
		Stock:                      in.Stock,                      // 库存
		LowStock:                   in.LowStock,                   // 库存预警值
		Unit:                       in.Unit,                       // 单位
		Weight:                     in.Weight,                     // 商品重量，默认为克
		PreviewStatus:              in.PreviewStatus,              // 是否为预告商品：0->不是；1->是
		ServiceIds:                 in.ServiceIds,                 // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
		Keywords:                   in.Keywords,                   // 搜索关键字
		Note:                       in.Note,                       // 备注
		AlbumPics:                  in.AlbumPics,                  // 画册图片，连产品图片限制为5张，以逗号分割
		DetailTitle:                in.DetailTitle,                // 详情标题
		DetailDesc:                 in.DetailDesc,                 // 详情描述
		DetailHTML:                 in.DetailHtml,                 // 产品详情网页内容
		DetailMobileHTML:           in.DetailMobileHtml,           // 移动端网页详情
		PromotionStartTime:         &PromotionStartTime,           // 促销开始时间
		PromotionEndTime:           &PromotionEndTime,             // 促销结束时间
		PromotionPerLimit:          in.PromotionPerLimit,          // 活动限购数量
		PromotionType:              in.PromotionType,              // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		BrandName:                  in.BrandName,                  // 品牌名称
		ProductCategoryName:        in.ProductCategoryName,        // 商品分类名称
		ProductCategoryIDArray:     in.ProductCategoryIdArray,     // 商品分类id字符串
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品失败")
	}

	// 根据促销类型设置价格：会员价格、阶梯价格、满减价格
	// 2.会员价格
	memberPrice := query.PmsMemberPrice.WithContext(l.ctx)
	_, _ = memberPrice.Where(query.PmsMemberPrice.ProductID.Eq(productId)).Delete()
	for _, list := range in.MemberPriceList {
		_ = memberPrice.Create(&model.PmsMemberPrice{
			ProductID:       productId,
			MemberLevelID:   list.MemberLevelId,
			MemberPrice:     list.MemberPrice,
			MemberLevelName: list.MemberLevelName,
		})
	}

	// 3.阶梯价格
	ladder := query.PmsProductLadder.WithContext(l.ctx)
	_, _ = ladder.Where(query.PmsProductLadder.ProductID.Eq(productId)).Delete()
	for _, list := range in.ProductLadderList {
		_ = ladder.Create(&model.PmsProductLadder{
			ProductID: productId,
			Count:     list.Count,
			Discount:  list.Discount,
			Price:     list.Price,
		})
	}
	// 4.满减价格
	full := query.PmsProductFullReduction.WithContext(l.ctx)
	_, _ = full.Where(query.PmsProductFullReduction.ProductID.Eq(productId)).Delete()
	for _, list := range in.ProductFullReductionList {
		_ = full.Create(&model.PmsProductFullReduction{
			ProductID:   productId,
			FullPrice:   list.FullPrice,
			ReducePrice: list.ReducePrice,
		})
	}

	// 5.更新sku库存信息
	sku := query.PmsSkuStock.WithContext(l.ctx)
	_, _ = sku.Where(query.PmsSkuStock.ProductID.Eq(productId)).Delete()
	for _, list := range in.SkuStockList {
		if list.SkuCode == "" {
			list.SkuCode = time.Now().Format("200601021504") + strconv.Itoa(rand.Intn(10))
		}
		_ = sku.Create(&model.PmsSkuStock{
			ProductID:      productId,
			SkuCode:        list.SkuCode,
			Price:          list.Price,
			Stock:          list.Stock,
			LowStock:       list.LowStock,
			Pic:            list.Pic,
			Sale:           list.Sale,
			PromotionPrice: list.PromotionPrice,
			LockStock:      list.LockStock,
			SpData:         list.SpData,
		})
	}
	// 6.更新商品参数,添加自定义商品规格
	attr := query.PmsProductAttributeValue.WithContext(l.ctx)
	_, _ = attr.Where(query.PmsProductAttributeValue.ProductID.Eq(productId)).Delete()
	for _, list := range in.ProductAttributeValueList {
		_ = attr.Create(&model.PmsProductAttributeValue{
			ProductID:          productId,
			ProductAttributeID: list.ProductAttributeId,
			Value:              list.AttributeValues,
		})
	}

	return &pmsclient.UpdateProductResp{}, nil
}
