package productcategoryattributerelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryAttributeRelationDeleteLogic 产品分类和属性关糸
/*
Author: LiuFeiHua
Date: 2024/5/8 9:57
*/
type ProductCategoryAttributeRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationDeleteLogic {
	return &ProductCategoryAttributeRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductCategoryAttributeRelationDelete 删除产品分类和属性关糸
func (l *ProductCategoryAttributeRelationDeleteLogic) ProductCategoryAttributeRelationDelete(in *pmsclient.ProductCategoryAttributeRelationDeleteReq) (*pmsclient.ProductCategoryAttributeRelationDeleteResp, error) {
	q := query.PmsProductCategoryAttributeRelation
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductCategoryAttributeRelationDeleteResp{}, nil
}
