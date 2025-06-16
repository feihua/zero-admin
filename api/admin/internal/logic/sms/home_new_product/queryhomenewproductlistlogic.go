package home_new_product

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeNewProductListLogic 首页新品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:27
*/
type HomeNewProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryHomeNewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductListLogic {
	return HomeNewProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeNewProductListLogic) QueryHomeNewProductList(req *types.QueryHomeNewProductListReq) (*types.QueryHomeNewProductListResp, error) {
	var resp, err = l.svcCtx.ProductSpuService.QueryProductSpuList(l.ctx, &pmsclient.QueryProductSpuListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		Name:            strings.TrimSpace(req.ProductName),
		VerifyStatus:    2,
		CategoryId:      0,
		PublishStatus:   2,
		BrandId:         0,
		RecommendStatus: 2,
		NewStatus:       1,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品信息列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询商品信息失败")
	}

	var list []*types.QueryProductSpuListData

	for _, detail := range resp.List {

		list = append(list, &types.QueryProductSpuListData{
			Id:                  detail.Id,                  // 商品SpuId
			Name:                detail.Name,                // 商品名称
			CategoryId:          detail.CategoryId,          // 商品分类ID
			CategoryIds:         detail.CategoryIds,         // 商品分类ID集合
			CategoryName:        detail.CategoryName,        // 商品分类名称
			BrandId:             detail.BrandId,             // 品牌ID
			BrandName:           detail.BrandName,           // 品牌名称
			Unit:                detail.Unit,                // 单位
			Weight:              detail.Weight,              // 重量(kg)
			Keywords:            detail.Keywords,            // 关键词
			Brief:               detail.Brief,               // 简介
			Description:         detail.Description,         // 详细描述
			AlbumPics:           detail.AlbumPics,           // 画册图片，最多8张，以逗号分割
			MainPic:             detail.MainPic,             // 主图
			PriceRange:          detail.PriceRange,          // 价格区间
			PublishStatus:       detail.PublishStatus,       // 上架状态：0-下架，1-上架
			NewStatus:           detail.NewStatus,           // 新品状态:0->不是新品；1->新品
			RecommendStatus:     detail.RecommendStatus,     // 推荐状态；0->不推荐；1->推荐
			VerifyStatus:        detail.VerifyStatus,        // 审核状态：0->未审核；1->审核通过
			PreviewStatus:       detail.PreviewStatus,       // 是否为预告商品：0->不是；1->是
			Sort:                detail.Sort,                // 排序
			NewStatusSort:       detail.NewStatusSort,       // 新品排序
			RecommendStatusSort: detail.RecommendStatusSort, // 推荐排序
			Sales:               detail.Sales,               // 销量
			Stock:               detail.Stock,               // 库存
			LowStock:            detail.LowStock,            // 预警库存
			PromotionType:       detail.PromotionType,       // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->秒杀
			DetailTitle:         detail.DetailTitle,         // 详情标题
			DetailDesc:          detail.DetailDesc,          // 详情描述
			DetailHtml:          detail.DetailHtml,          // 产品详情网页内容
			DetailMobileHtml:    detail.DetailMobileHtml,    // 移动端网页详情
			CreateBy:            detail.CreateBy,            // 创建人ID
			CreateTime:          detail.CreateTime,          // 创建时间
			UpdateBy:            detail.UpdateBy,            // 更新人ID
			UpdateTime:          detail.UpdateTime,          // 更新时间
		})
	}

	return &types.QueryHomeNewProductListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询新鲜好物表成功",
	}, nil
}
