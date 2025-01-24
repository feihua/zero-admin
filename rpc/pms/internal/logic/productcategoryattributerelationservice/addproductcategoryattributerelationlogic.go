package productcategoryattributerelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductCategoryAttributeRelationLogic 添加产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
/*
Author: LiuFeiHua
Date: 2024/6/12 16:51
*/
type AddProductCategoryAttributeRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductCategoryAttributeRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductCategoryAttributeRelationLogic {
	return &AddProductCategoryAttributeRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductCategoryAttributeRelation 添加产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
func (l *AddProductCategoryAttributeRelationLogic) AddProductCategoryAttributeRelation(in *pmsclient.AddProductCategoryAttributeRelationReq) (*pmsclient.AddProductCategoryAttributeRelationResp, error) {
	err := query.PmsProductCategoryAttributeRelation.WithContext(l.ctx).Create(&model.PmsProductCategoryAttributeRelation{
		ProductCategoryID:  in.ProductCategoryId,  // 商品分类id
		ProductAttributeID: in.ProductAttributeId, // 商品属性id
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AddProductCategoryAttributeRelationResp{}, nil
}
