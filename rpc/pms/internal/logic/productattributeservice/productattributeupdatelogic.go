package productattributeservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeUpdateLogic 商品属性
/*
Author: LiuFeiHua
Date: 2024/5/8 10:52
*/
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

// ProductAttributeUpdate 更新商品属性
func (l *ProductAttributeUpdateLogic) ProductAttributeUpdate(in *pmsclient.ProductAttributeUpdateReq) (*pmsclient.ProductAttributeUpdateResp, error) {
	q := query.PmsProductAttribute
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductAttribute{
		ID:                         in.Id,
		ProductAttributeCategoryID: in.ProductAttributeCategoryId,
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

	return &pmsclient.ProductAttributeUpdateResp{}, nil
}
