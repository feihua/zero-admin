package category

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryUpdateLogic 商品分类
/*
Author: LiuFeiHua
Date: 2024/5/15 11:17
*/
type ProductCategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryUpdateLogic {
	return ProductCategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductCategoryUpdate 更新商品分类
func (l *ProductCategoryUpdateLogic) ProductCategoryUpdate(req types.UpdateProductCategoryReq) (*types.UpdateProductCategoryResp, error) {
	_, err := l.svcCtx.ProductCategoryService.UpdateProductCategory(l.ctx, &pmsclient.UpdateProductCategoryReq{
		Id:                     req.Id,
		ParentId:               req.ParentId,
		Name:                   req.Name,
		Level:                  req.Level,
		ProductCount:           req.ProductCount,
		ProductUnit:            req.ProductUnit,
		NavStatus:              req.NavStatus,
		ShowStatus:             req.ShowStatus,
		Sort:                   req.Sort,
		Icon:                   req.Icon,
		Keywords:               req.Keywords,
		Description:            req.Description,
		ProductAttributeIdList: req.ProductAttributeIdList,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品分类信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新商品分类失败")
	}

	return &types.UpdateProductCategoryResp{
		Code:    "000000",
		Message: "更新商品分类成功",
	}, nil
}
