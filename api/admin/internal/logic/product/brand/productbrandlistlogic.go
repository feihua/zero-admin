package brand

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

// ProductBrandListLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/13 16:54
*/
type ProductBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductBrandListLogic {
	return ProductBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductBrandList 查询商品品牌
func (l *ProductBrandListLogic) ProductBrandList(req types.ListProductBrandReq) (*types.ListProductBrandResp, error) {
	resp, err := l.svcCtx.BrandService.QueryBrandList(l.ctx, &pmsclient.QueryBrandListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		Name:            strings.TrimSpace(req.Name),
		FactoryStatus:   req.FactoryStatus,
		ShowStatus:      req.ShowStatus,
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品品牌列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询商品品牌失败")
	}

	var list []*types.ListProductBrandData

	for _, item := range resp.List {
		list = append(list, &types.ListProductBrandData{
			Id:                  item.Id,                  //
			Name:                item.Name,                // 品牌名称
			FirstLetter:         item.FirstLetter,         // 首字母
			Sort:                item.Sort,                // 排序
			FactoryStatus:       item.FactoryStatus,       // 是否为品牌制造商：0->不是；1->是
			ShowStatus:          item.ShowStatus,          // 品牌显示状态
			RecommendStatus:     item.RecommendStatus,     // 推荐状态
			ProductCount:        item.ProductCount,        // 产品数量
			ProductCommentCount: item.ProductCommentCount, // 产品评论数量
			Logo:                item.Logo,                // 品牌logo
			BigPic:              item.BigPic,              // 专区大图
			BrandStory:          item.BrandStory,          // 品牌故事
			CreateBy:            item.CreateBy,            // 创建者
			CreateTime:          item.CreateTime,          // 创建时间
			UpdateBy:            item.UpdateBy,            // 更新者
			UpdateTime:          item.UpdateTime,          // 更新时间
		})
	}

	return &types.ListProductBrandResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询商品品牌成功",
	}, nil
}
