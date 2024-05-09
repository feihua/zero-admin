package productcategoryattributerelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryAttributeRelationAddLogic 产品分类和属性关糸
/*
Author: LiuFeiHua
Date: 2024/5/8 10:56
*/
type ProductCategoryAttributeRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationAddLogic {
	return &ProductCategoryAttributeRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductCategoryAttributeRelationAdd 添加产品分类和属性关糸
func (l *ProductCategoryAttributeRelationAddLogic) ProductCategoryAttributeRelationAdd(in *pmsclient.ProductCategoryAttributeRelationAddReq) (*pmsclient.ProductCategoryAttributeRelationAddResp, error) {
	err := query.PmsProductCategoryAttributeRelation.WithContext(l.ctx).Create(&model.PmsProductCategoryAttributeRelation{
		ProductCategoryID:  in.ProductCategoryId,
		ProductAttributeID: in.ProductAttributeId,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductCategoryAttributeRelationAddResp{}, nil
}
