package attributecategory

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributecategoryAddLogic 商品属性分类
/*
Author: LiuFeiHua
Date: 2024/5/15 13:53
*/
type ProductAttributecategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributecategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributecategoryAddLogic {
	return &ProductAttributecategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductAttributecategoryAdd 添加商品属性分类
func (l *ProductAttributecategoryAddLogic) ProductAttributecategoryAdd(req *types.AddProductAttributecategoryReq) (resp *types.AddProductAttributecategoryResp, err error) {
	_, err = l.svcCtx.ProductAttributeCategoryService.ProductAttributeCategoryAdd(l.ctx, &pmsclient.ProductAttributeCategoryAddReq{
		Name: req.Name,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加属性分类信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加属性分类失败")
	}

	return &types.AddProductAttributecategoryResp{
		Code:    "000000",
		Message: "添加属性分类成功",
	}, nil
}
