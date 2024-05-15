package productattributeservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeDeleteLogic 商品属性
/*
Author: LiuFeiHua
Date: 2024/5/8 9:54
*/
type ProductAttributeDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeDeleteLogic {
	return &ProductAttributeDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAttributeDelete 删除商品属性
// 1.根据属性id查询属性信息
// 2.根据属性分类id查询分类信息
func (l *ProductAttributeDeleteLogic) ProductAttributeDelete(in *pmsclient.ProductAttributeDeleteReq) (*pmsclient.ProductAttributeDeleteResp, error) {
	q := query.PmsProductAttribute
	category := query.PmsProductAttributeCategory
	for _, id := range in.Ids {
		//1.根据属性id查询属性信息
		attribute, _ := q.WithContext(l.ctx).Where(q.ID.Eq(id)).First()
		//2.根据属性分类id查询分类信息
		categoryDo := category.WithContext(l.ctx).Where(category.ID.Eq(attribute.ProductAttributeCategoryID))
		first, _ := categoryDo.First()

		if attribute.Type == 0 {
			_, _ = categoryDo.Update(category.AttributeCount, first.AttributeCount-1)
		}

		if attribute.Type == 1 {
			_, _ = categoryDo.Update(category.ParamCount, first.ParamCount-1)
		}
	}

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeDeleteResp{}, nil
}
