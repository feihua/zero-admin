package productattributeservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeAddLogic 商品属性
/*
Author: LiuFeiHua
Date: 2024/5/8 10:51
*/
type ProductAttributeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeAddLogic {
	return &ProductAttributeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAttributeAdd 添加商品属性
func (l *ProductAttributeAddLogic) ProductAttributeAdd(in *pmsclient.ProductAttributeAddReq) (*pmsclient.ProductAttributeAddResp, error) {
	err := query.PmsProductAttribute.WithContext(l.ctx).Create(&model.PmsProductAttribute{
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

	return &pmsclient.ProductAttributeAddResp{}, nil
}
