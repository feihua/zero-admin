package productcategoryattributerelationservicelogic

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ProductCategoryAttributeRelationDeleteLogic) ProductCategoryAttributeRelationDelete(in *pmsclient.ProductCategoryAttributeRelationDeleteReq) (*pmsclient.ProductCategoryAttributeRelationDeleteResp, error) {
	err := l.svcCtx.PmsProductCategoryAttributeRelationModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductCategoryAttributeRelationDeleteResp{}, nil
}
