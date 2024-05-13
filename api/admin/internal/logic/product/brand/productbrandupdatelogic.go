package brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductBrandUpdateLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/13 16:55
*/
type ProductBrandUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductBrandUpdateLogic {
	return ProductBrandUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductBrandUpdate 更新商品品牌
func (l *ProductBrandUpdateLogic) ProductBrandUpdate(req types.UpdateProductBrandReq) (*types.UpdateProductBrandResp, error) {
	_, err := l.svcCtx.BrandService.BrandUpdate(l.ctx, &pmsclient.BrandUpdateReq{
		Id:                  req.Id,
		Name:                req.Name,
		FirstLetter:         req.FirstLetter,
		Sort:                req.Sort,
		FactoryStatus:       req.FactoryStatus,
		ShowStatus:          req.ShowStatus,
		ProductCount:        req.ProductCount,
		ProductCommentCount: req.ProductCommentCount,
		Logo:                req.Logo,
		BigPic:              req.BigPic,
		BrandStory:          req.BrandStory,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品品牌信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新商品品牌失败")
	}

	return &types.UpdateProductBrandResp{
		Code:    "000000",
		Message: "更新商品品牌成功",
	}, nil
}
