package productspuservicelogic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

// UpdateProductSpuLogic 更新商品SPU
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:38
*/
type UpdateProductSpuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSpuLogic {
	return &UpdateProductSpuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductSpu 更新商品SPU
// 逻辑步骤：
// 1.更新商品基本信息
// 2.会员价格
// 3.阶梯价格
// 4.满减价格
// 5.更新sku库存信息
// 6.更新商品参数,添加自定义商品规格
func (l *UpdateProductSpuLogic) UpdateProductSpu(in *pmsclient.ProductSpuReq) (*pmsclient.ProductSpuResp, error) {
	spu := query.PmsProductSpu
	q := spu.WithContext(l.ctx)

	// 1.根据商品SPUid查询商品SPU是否已存在
	detail, err := q.Where(spu.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品SPU不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品SPU不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品SPU异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品SPU异常")
	}

	now := time.Now()
	item := &model.PmsProductSpu{
		ID:                  in.Id,                  // 商品SpuId
		ProductSn:           in.ProductSn,           // 商品货号
		Name:                in.Name,                // 商品名称
		CategoryID:          in.CategoryId,          // 商品分类ID
		CategoryIds:         in.CategoryIds,         // 商品分类ID集合
		CategoryName:        in.CategoryName,        // 商品分类名称
		BrandID:             in.BrandId,             // 品牌ID
		BrandName:           in.BrandName,           // 品牌名称
		Unit:                in.Unit,                // 单位
		Weight:              float64(in.Weight),     // 重量(kg)
		Keywords:            in.Keywords,            // 关键词
		Brief:               in.Brief,               // 简介
		Description:         in.Description,         // 详细描述
		AlbumPics:           in.AlbumPics,           // 画册图片，最多8张，以逗号分割
		MainPic:             in.MainPic,             // 主图
		PriceRange:          in.PriceRange,          // 价格区间
		PublishStatus:       in.PublishStatus,       // 上架状态：0-下架，1-上架
		NewStatus:           in.NewStatus,           // 新品状态:0->不是新品；1->新品
		RecommendStatus:     in.RecommendStatus,     // 推荐状态；0->不推荐；1->推荐
		VerifyStatus:        in.VerifyStatus,        // 审核状态：0->未审核；1->审核通过
		PreviewStatus:       in.PreviewStatus,       // 是否为预告商品：0->不是；1->是
		Sort:                in.Sort,                // 排序
		NewStatusSort:       in.NewStatusSort,       // 新品排序
		RecommendStatusSort: in.RecommendStatusSort, // 推荐排序
		Sales:               in.Sales,               // 销量
		Stock:               in.Stock,               // 库存
		LowStock:            in.LowStock,            // 预警库存
		PromotionType:       in.PromotionType,       // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀
		DetailTitle:         in.DetailTitle,         // 详情标题
		DetailDesc:          in.DetailDesc,          // 详情描述
		DetailHTML:          in.DetailHtml,          // 产品详情网页内容
		DetailMobileHTML:    in.DetailMobileHtml,    // 移动端网页详情
		CreateBy:            detail.CreateBy,        // 创建人ID
		CreateTime:          detail.CreateTime,      // 创建时间
		UpdateBy:            &in.CreateBy,           // 更新人ID
		UpdateTime:          &now,                   // 更新时间
	}

	// 2.商品SPU存在时,则直接更新商品SPU
	err = l.svcCtx.DB.Model(&model.PmsProductSpu{}).WithContext(l.ctx).Where(spu.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新商品SPU失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新商品SPU失败")
	}

	spuId := in.Id
	// 根据促销类型设置价格：会员价格、阶梯价格、满减价格
	// 2.会员价格
	memberPrice := query.PmsMemberPrice.WithContext(l.ctx)
	_, _ = memberPrice.Where(query.PmsMemberPrice.ProductID.Eq(spuId)).Delete()
	for _, list := range in.MemberPriceList {
		_ = memberPrice.Create(&model.PmsMemberPrice{
			ProductID:       spuId,
			MemberLevelID:   list.LevelId,
			MemberPrice:     list.Price,
			MemberLevelName: list.LevelName,
		})
	}

	// 3.阶梯价格
	ladder := query.PmsProductLadder.WithContext(l.ctx)
	_, _ = ladder.Where(query.PmsProductLadder.ProductID.Eq(spuId)).Delete()
	for _, list := range in.ProductLadderList {
		_ = ladder.Create(&model.PmsProductLadder{
			ProductID: spuId,
			Count:     list.Count,
			Discount:  list.Discount,
			Price:     list.Price,
		})
	}
	// 4.满减价格
	full := query.PmsProductFullReduction.WithContext(l.ctx)
	_, _ = full.Where(query.PmsProductFullReduction.ProductID.Eq(spuId)).Delete()
	for _, list := range in.ProductFullReductionList {
		_ = full.Create(&model.PmsProductFullReduction{
			ProductID:   spuId,
			FullPrice:   list.FullPrice,
			ReducePrice: list.ReducePrice,
		})
	}

	// 5.更新sku库存信息
	sku := query.PmsProductSku.WithContext(l.ctx)
	_, _ = sku.Where(query.PmsProductSku.SpuID.Eq(spuId)).Delete()
	for _, list := range in.SkuStockList {
		skuCode := time.Now().Format("200601021504") + strconv.Itoa(rand.Intn(10))
		_ = sku.Create(&model.PmsProductSku{
			SpuID:          spuId,                        // 商品SpuId
			Name:           list.Name,                    // SKU名称
			SkuCode:        skuCode,                      // SKU编码
			MainPic:        list.MainPic,                 // 主图
			AlbumPics:      list.AlbumPics,               // 图片集
			Price:          float64(list.Price),          // 价格
			PromotionPrice: float64(list.PromotionPrice), // 单品促销价格
			Stock:          list.Stock,                   // 库存
			LowStock:       list.LowStock,                // 预警库存
			SpecData:       list.SpecData,                // 规格数据
			Weight:         float64(list.Weight),         // 重量(kg)
			PublishStatus:  list.PublishStatus,           // 上架状态：0-下架，1-上架
			VerifyStatus:   list.VerifyStatus,            // 审核状态：0-未审核，1-审核通过，2-审核不通过
			Sort:           list.Sort,                    // 排序
			CreateBy:       in.CreateBy,                  // 创建人ID
		})
	}
	// 6.更新商品参数,添加自定义商品规格
	attr := query.PmsProductAttributeValue.WithContext(l.ctx)
	_, _ = attr.Where(query.PmsProductAttributeValue.SpuID.Eq(spuId)).Delete()
	for _, list := range in.ProductAttributeValueList {
		_ = attr.Create(&model.PmsProductAttributeValue{
			SpuID:       spuId,                   // 商品SPU ID
			AttributeID: list.ProductAttributeId, // 属性ID
			Value:       list.AttributeValues,    // 属性值
			CreateBy:    in.CreateBy,             // 创建人ID
		})
	}

	message := map[string]any{"id": spuId}
	body, _ := json.Marshal(message)
	err = l.svcCtx.RabbitMQ.SendMessage("product.event.exchange", "syn.product.to.es.queue", "syn.product.key", body)

	return &pmsclient.ProductSpuResp{}, nil
}
