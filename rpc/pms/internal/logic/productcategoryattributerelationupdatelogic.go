package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ProductCategoryAttributeRelationUpdateLogic) ProductCategoryAttributeRelationUpdate(in *pms.ProductCategoryAttributeRelationUpdateReq) (*pms.ProductCategoryAttributeRelationUpdateResp, error) {
	err := l.svcCtx.PmsProductCategoryAttributeRelationModel.Update(pmsmodel.PmsProductCategoryAttributeRelation{
		Id:                 in.Id,
		ProductCategoryId:  in.ProductCategoryId,
		ProductAttributeId: in.ProductAttributeId,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductCategoryAttributeRelationUpdateResp{}, nil
}
