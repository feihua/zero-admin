package productcategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryAddLogic 商品类别
/*
Author: LiuFeiHua
Date: 2024/5/8 10:59
*/
type ProductCategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAddLogic {
	return &ProductCategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductCategoryAdd 添加商品类别
func (l *ProductCategoryAddLogic) ProductCategoryAdd(in *pmsclient.ProductCategoryAddReq) (*pmsclient.ProductCategoryAddResp, error) {
	category := &model.PmsProductCategory{
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
	}
	err := query.PmsProductCategory.WithContext(l.ctx).Create(category)

	if err != nil {
		return nil, err
	}

	if len(in.ProductAttributeIdList) > 0 {
		var list []*model.PmsProductCategoryAttributeRelation
		for _, productAttributeId := range in.ProductAttributeIdList {
			list = append(list, &model.PmsProductCategoryAttributeRelation{
				ProductCategoryID:  category.ID,
				ProductAttributeID: productAttributeId,
			})
		}

		_ = query.PmsProductCategoryAttributeRelation.WithContext(l.ctx).CreateInBatches(list, len(list))
	}

	return &pmsclient.ProductCategoryAddResp{}, nil
}
