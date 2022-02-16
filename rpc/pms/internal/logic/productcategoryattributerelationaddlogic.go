package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ProductCategoryAttributeRelationAddLogic) ProductCategoryAttributeRelationAdd(in *pms.ProductCategoryAttributeRelationAddReq) (*pms.ProductCategoryAttributeRelationAddResp, error) {
	_, err := l.svcCtx.PmsProductCategoryAttributeRelationModel.Insert(pmsmodel.PmsProductCategoryAttributeRelation{
		ProductCategoryId:  in.ProductCategoryId,
		ProductAttributeId: in.ProductAttributeId,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductCategoryAttributeRelationAddResp{}, nil
}
