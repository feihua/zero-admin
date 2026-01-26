// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package brand

import (
	"context"
	"strings"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryBrandDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryBrandDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBrandDetailLogic {
	return &QueryBrandDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryBrandDetailLogic) QueryBrandDetail(req *types.QueryBrandDetailReq) (resp *types.QueryBrandDetailResp, err error) {
	detail, err := l.svcCtx.ProductBrandService.QueryProductBrandDetail(l.ctx, &pmsclient.QueryProductBrandDetailReq{
		Id: req.BrandId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品品牌详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	brandData := types.BrandData{
		Id:                  detail.Id,                  //
		Name:                detail.Name,                // 品牌名称
		Logo:                detail.Logo,                // 品牌logo
		BigPic:              detail.BigPic,              // 专区大图
		Description:         detail.Description,         // 描述
		FirstLetter:         detail.FirstLetter,         // 首字母
		Sort:                detail.Sort,                // 排序
		RecommendStatus:     detail.RecommendStatus,     // 推荐状态
		ProductCount:        detail.ProductCount,        // 产品数量
		ProductCommentCount: detail.ProductCommentCount, // 产品评论数量

	}

	productListResp, err := l.svcCtx.ProductSpuService.QueryProductSpuList(l.ctx, &pmsclient.QueryProductSpuListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		CategoryId:      0,           // 商品分类ID
		BrandId:         req.BrandId, // 品牌ID
		PublishStatus:   1,           // 上架状态：0-下架，1-上架
		NewStatus:       2,           // 新品状态:0->不是新品；1->新品
		RecommendStatus: 2,           // 推荐状态；0->不推荐；1->推荐
		VerifyStatus:    1,           // 审核状态：0->未审核；1->审核通过
		PreviewStatus:   0,           // 是否为预告商品：0->不是；1->是
		PromotionType:   6,           // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀

	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品列表失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	productLists := make([]types.BrandProductData, 0)
	for _, product := range productListResp.List {
		price := strings.Split(product.PriceRange, "-")[0]
		productLists = append(productLists, types.BrandProductData{
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
			AlbumPics:           product.AlbumPics,           // 画册图片，最多8张，以逗号分割
			MainPic:             product.MainPic,             // 主图
			Price:               price,                       // 价格
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
			SubTitle:            product.SubTitle,            // 详情标题
			DetailHtml:          product.DetailHtml,          // 产品详情网页内容
			DetailMobileHtml:    product.DetailMobileHtml,    // 移动端网页详情
		})
	}

	return &types.QueryBrandDetailResp{
		Code:    0,
		Message: "操作成功",
		Data: types.BrandDetailData{
			BrandData:        brandData,
			BrandProductData: productLists,
		},
	}, nil
}
