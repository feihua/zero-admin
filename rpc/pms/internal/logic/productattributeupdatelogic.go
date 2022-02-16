package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeUpdateLogic {
	return &ProductAttributeUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeUpdateLogic) ProductAttributeUpdate(in *pms.ProductAttributeUpdateReq) (*pms.ProductAttributeUpdateResp, error) {
	err := l.svcCtx.PmsProductAttributeModel.Update(pmsmodel.PmsProductAttribute{
		Id:                         in.Id,
		ProductAttributeCategoryId: in.ProductAttributeCategoryId,
		Name:                       in.Name,
		SelectType:                 in.SelectType,
		InputType:                  in.InputType,
		InputList:                  in.InputList,
		Sort:                       in.Sort,
		FilterType:                 in.FilterType,
		SearchType:                 in.SearchType,
		RelatedStatus:              in.RelatedStatus,
		HandAddStatus:              in.HandAddStatus,
		Type:                       in.Type,
	})
	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeUpdateResp{}, nil
}
