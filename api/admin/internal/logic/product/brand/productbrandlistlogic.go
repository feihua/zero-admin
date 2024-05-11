package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ProductBrandListLogic) ProductBrandList(req types.ListProductBrandReq) (*types.ListProductBrandResp, error) {
	resp, err := l.svcCtx.BrandService.BrandList(l.ctx, &pmsclient.BrandListReq{
		Current:       req.Current,
		PageSize:      req.PageSize,
		Name:          req.Name,
		FactoryStatus: req.FactoryStatus,
		ShowStatus:    req.ShowStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品品牌列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询商品品牌失败")
	}

	var list []*types.ListtProductBrandData

	for _, item := range resp.List {
		list = append(list, &types.ListtProductBrandData{
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
