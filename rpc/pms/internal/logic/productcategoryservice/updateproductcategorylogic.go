package productcategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductCategoryLogic 更新产品分类
/*
Author: LiuFeiHua
Date: 2024/6/12 17:01
*/
type UpdateProductCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductCategoryLogic {
	return &UpdateProductCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductCategory 更新产品分类
func (l *UpdateProductCategoryLogic) UpdateProductCategory(in *pmsclient.UpdateProductCategoryReq) (*pmsclient.UpdateProductCategoryResp, error) {
	//更新商品分类时要更新商品中的名称
	product := query.PmsProduct
	_, _ = product.WithContext(l.ctx).Where(product.ProductCategoryID.Eq(in.Id)).Update(product.ProductCategoryName, in.Name)

	//同时更新筛选属性的信息
	relation := query.PmsProductCategoryAttributeRelation
	_, _ = relation.WithContext(l.ctx).Where(relation.ProductCategoryID.Eq(in.Id)).Delete()
	if len(in.ProductAttributeIdList) > 0 {
		var list []*model.PmsProductCategoryAttributeRelation
		for _, productAttributeId := range in.ProductAttributeIdList {
			list = append(list, &model.PmsProductCategoryAttributeRelation{
				ProductCategoryID:  in.Id,
				ProductAttributeID: productAttributeId,
			})
		}
		_ = relation.WithContext(l.ctx).CreateInBatches(list, len(list))
	}

	q := query.PmsProductCategory
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductCategory{
		ID:           in.Id,
		ParentID:     in.ParentId,
		Name:         in.Name,
		Level:        in.Level,
		ProductCount: in.ProductCount,
		ProductUnit:  in.ProductUnit,
		NavStatus:    in.NavStatus,
		ShowStatus:   in.ShowStatus,
		Sort:         in.Sort,
		Icon:         in.Icon,
		Keywords:     in.Keywords,
		Description:  &in.Description,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.UpdateProductCategoryResp{}, nil
}
