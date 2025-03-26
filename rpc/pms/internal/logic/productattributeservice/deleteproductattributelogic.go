package productattributeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductAttributeLogic 删除商品属性参数
/*
Author: LiuFeiHua
Date: 2024/6/12 16:48
*/
type DeleteProductAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductAttributeLogic {
	return &DeleteProductAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductAttribute 删除商品属性参数
// 1.根据属性id查询属性信息
// 2.根据属性分类id查询分类信息
func (l *DeleteProductAttributeLogic) DeleteProductAttribute(in *pmsclient.DeleteProductAttributeReq) (*pmsclient.DeleteProductAttributeResp, error) {
	q := query.PmsProductAttribute
	category := query.PmsProductAttributeCategory
	for _, id := range in.Ids {
		// 1.根据属性id查询属性信息
		attribute, _ := q.WithContext(l.ctx).Where(q.ID.Eq(id)).First()
		// 2.根据属性分类id查询分类信息
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
		logc.Errorf(l.ctx, "删除商品属性参数失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品属性参数失败")
	}

	return &pmsclient.DeleteProductAttributeResp{}, nil
}
