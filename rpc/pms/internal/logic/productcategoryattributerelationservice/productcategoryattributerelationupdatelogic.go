package productcategoryattributerelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryAttributeRelationUpdateLogic 产品分类和属性关糸
/*
Author: LiuFeiHua
Date: 2024/5/8 10:57
*/
type ProductCategoryAttributeRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationUpdateLogic {
	return &ProductCategoryAttributeRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductCategoryAttributeRelationUpdate 更新产品分类和属性关糸
func (l *ProductCategoryAttributeRelationUpdateLogic) ProductCategoryAttributeRelationUpdate(in *pmsclient.ProductCategoryAttributeRelationUpdateReq) (*pmsclient.ProductCategoryAttributeRelationUpdateResp, error) {
	q := query.PmsProductCategoryAttributeRelation
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductCategoryAttributeRelation{
		ID:                 in.Id,
		ProductCategoryID:  in.ProductCategoryId,
		ProductAttributeID: in.ProductAttributeId,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductCategoryAttributeRelationUpdateResp{}, nil
}
