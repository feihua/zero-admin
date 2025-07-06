package product

import (
	"context"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductListLogic {
	return &QueryProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryProductListLogic) QueryProductList(req *types.QueryProductListReq) (resp *types.QueryProductListResp, err error) {
	productListResp, err := l.svcCtx.ProductSpuService.QueryProductSpuList(l.ctx, &pmsclient.QueryProductSpuListReq{
		PageNum:       req.Current,
		PageSize:      req.PageSize,
		VerifyStatus:  1, // 审核状态：0->未审核；1->审核通过
		CategoryId:    req.ProductCategoryId,
		PublishStatus: 1, // 上架状态：0->下架；1->上架
		BrandId:       req.BrandId,
		Name:          req.Keyword,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品列表失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var productLists []types.ProductData
	for _, product := range productListResp.List {
		productLists = append(productLists, types.ProductData{
			Id:                  product.Id,                  // 商品SpuId
			Name:                product.Name,                // 商品名称
			ProductSn:           product.ProductSn,           // 商品货号
			CategoryId:          product.CategoryId,          // 商品分类ID
			CategoryIds:         product.CategoryIds,         // 商品分类ID集合
			CategoryName:        product.CategoryName,        // 商品分类名称
			BrandId:             product.BrandId,             // 品牌ID
			BrandName:           product.BrandName,           // 品牌名称
			Unit:                product.Unit,                // 单位
			Weight:              product.Weight,              // 重量(kg)
			Keywords:            product.Keywords,            // 关键词
			Brief:               product.Brief,               // 简介
			Description:         product.Description,         // 详细描述
			AlbumPics:           product.AlbumPics,           // 画册图片，最多8张，以逗号分割
			MainPic:             product.MainPic,             // 主图
			PriceRange:          product.PriceRange,          // 价格区间
			PublishStatus:       product.PublishStatus,       // 上架状态：0-下架，1-上架
			NewStatus:           product.NewStatus,           // 新品状态:0->不是新品；1->新品
			RecommendStatus:     product.RecommendStatus,     // 推荐状态；0->不推荐；1->推荐
			VerifyStatus:        product.VerifyStatus,        // 审核状态：0->未审核；1->审核通过
			PreviewStatus:       product.PreviewStatus,       // 是否为预告商品：0->不是；1->是
			Sort:                product.Sort,                // 排序
			NewStatusSort:       product.NewStatusSort,       // 新品排序
			RecommendStatusSort: product.RecommendStatusSort, // 推荐排序
			Sales:               product.Sales,               // 销量
			Stock:               product.Stock,               // 库存
			LowStock:            product.LowStock,            // 预警库存
			PromotionType:       product.PromotionType,       // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀
			DetailTitle:         product.DetailTitle,         // 详情标题
			DetailDesc:          product.DetailDesc,          // 详情描述
			DetailHtml:          product.DetailHtml,          // 产品详情网页内容
			DetailMobileHtml:    product.DetailMobileHtml,    // 移动端网页详情
		})
	}

	return &types.QueryProductListResp{
		Code:    0,
		Message: "操作成功",
		Data:    productLists,
	}, nil
}
