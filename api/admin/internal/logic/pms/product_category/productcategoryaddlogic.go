package product_category

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

// ProductCategoryAddLogic 商品分类
/*
Author: LiuFeiHua
Date: 2024/5/15 11:16
*/
type ProductCategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryAddLogic {
	return ProductCategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductCategoryAdd 添加商品分类
func (l *ProductCategoryAddLogic) ProductCategoryAdd(req *types.AddProductCategoryReq) (*types.BaseResp, error) {
	_, err := l.svcCtx.ProductCategoryService.AddProductCategory(l.ctx, &pmsclient.AddProductCategoryReq{
		ParentId:     req.ParentId,    // 上机分类的编号：0表示一级分类
		Name:         req.Name,        // 商品分类名称
		Level:        req.Level,       // 分类级别：0->1级；1->2级
		ProductCount: 0,               // 商品数量
		ProductUnit:  req.ProductUnit, // 商品数量
		NavStatus:    req.NavStatus,   // 是否显示在导航栏：0->不显示；1->显示
		ShowStatus:   req.ShowStatus,  // 显示状态：0->不显示；1->显示
		Sort:         req.Sort,        // 排序
		Icon:         req.Icon,        // 图标
		Keywords:     req.Keywords,    // 关键字
		Description:  req.Description, // 描述
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品分类信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加商品分类失败")
	}

	return res.Success()
}
