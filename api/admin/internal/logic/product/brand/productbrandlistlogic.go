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
		PageNum:       req.Current,
		PageSize:      req.PageSize,
		Name:          strings.TrimSpace(req.Name),
		FactoryStatus: req.FactoryStatus,
		ShowStatus:    req.ShowStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品品牌列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询商品品牌失败")
	}

	var list []*types.ListProductBrandData

	for _, item := range resp.List {
		list = append(list, &types.ListProductBrandData{
			Id:                  item.Id,
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
