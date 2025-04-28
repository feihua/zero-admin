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

// ProductCategoryDeleteLogic 商品分类
/*
Author: LiuFeiHua
Date: 2024/5/15 11:16
*/
type ProductCategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryDeleteLogic {
	return ProductCategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductCategoryDelete 删除商品分类
func (l *ProductCategoryDeleteLogic) ProductCategoryDelete(req *types.DeleteProductCategoryReq) (*types.BaseResp, error) {
	_, err := l.svcCtx.ProductCategoryService.DeleteProductCategory(l.ctx, &pmsclient.DeleteProductCategoryReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除商品类别异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除商品类别失败")
	}

	return res.Success()
}
