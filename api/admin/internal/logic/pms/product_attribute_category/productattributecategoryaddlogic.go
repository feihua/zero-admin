package product_attribute_category

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeCategoryAddLogic 商品属性分类
/*
Author: LiuFeiHua
Date: 2024/5/15 13:53
*/
type ProductAttributeCategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributeCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryAddLogic {
	return &ProductAttributeCategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductAttributeCategoryAdd 添加商品属性分类
func (l *ProductAttributeCategoryAddLogic) ProductAttributeCategoryAdd(req *types.AddProductAttributeCategoryReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductAttributeCategoryService.AddProductAttributeCategory(l.ctx, &pmsclient.AddProductAttributeCategoryReq{
		Name: req.Name,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加属性分类信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加属性分类失败")
	}

	return res.Success()
}
