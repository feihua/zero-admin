package product_attribute_category

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeCategoryUpdateLogic 商品属性分类
type ProductAttributeCategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributeCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryUpdateLogic {
	return &ProductAttributeCategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductAttributeCategoryUpdate 更新商品属性分类
func (l *ProductAttributeCategoryUpdateLogic) ProductAttributeCategoryUpdate(req *types.UpdateProductAttributeCategoryReq) (resp *types.UpdateProductAttributeCategoryResp, err error) {
	_, err = l.svcCtx.ProductAttributeCategoryService.UpdateProductAttributeCategory(l.ctx, &pmsclient.UpdateProductAttributeCategoryReq{
		Id:   req.Id,
		Name: req.Name,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新属性分类信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新属性分类失败")
	}

	return &types.UpdateProductAttributeCategoryResp{
		Code:    "000000",
		Message: "更新属性分类成功",
	}, nil
}
