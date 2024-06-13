package productattributeservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductAttributeLogic 添加商品属性参数表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:47
*/
type AddProductAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductAttributeLogic {
	return &AddProductAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductAttribute 添加商品属性参数表
func (l *AddProductAttributeLogic) AddProductAttribute(in *pmsclient.AddProductAttributeReq) (*pmsclient.AddProductAttributeResp, error) {
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

	//新增商品属性以后需要更新商品属性分类数量
	q := query.PmsProductAttributeCategory
	categoryDo := q.WithContext(l.ctx).Where(q.ID.Eq(in.ProductAttributeCategoryId))
	category, err := categoryDo.First()
	if err != nil {
		return nil, err
	}

	if in.Type == 0 {
		_, _ = categoryDo.Update(q.AttributeCount, category.AttributeCount+1)
	}

	if in.Type == 1 {
		_, _ = categoryDo.Update(q.ParamCount, category.ParamCount+1)
	}

	return &pmsclient.AddProductAttributeResp{}, nil
}
