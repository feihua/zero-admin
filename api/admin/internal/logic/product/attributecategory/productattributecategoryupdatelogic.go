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

// ProductAttributecategoryUpdateLogic 商品属性分类
type ProductAttributecategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributecategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributecategoryUpdateLogic {
	return &ProductAttributecategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductAttributecategoryUpdate 更新商品属性分类
func (l *ProductAttributecategoryUpdateLogic) ProductAttributecategoryUpdate(req *types.UpdateProductAttributecategoryReq) (resp *types.UpdateProductAttributecategoryResp, err error) {
	_, err = l.svcCtx.ProductAttributeCategoryService.UpdateProductAttributeCategory(l.ctx, &pmsclient.UpdateProductAttributeCategoryReq{
		Id:   req.Id,
		Name: req.Name,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新属性分类信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新属性分类失败")
	}

	return &types.UpdateProductAttributecategoryResp{
		Code:    "000000",
		Message: "更新属性分类成功",
	}, nil
}
