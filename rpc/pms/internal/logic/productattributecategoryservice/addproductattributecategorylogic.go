package productattributecategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductAttributeCategoryLogic 添加产品属性分类表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:42
*/
type AddProductAttributeCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductAttributeCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductAttributeCategoryLogic {
	return &AddProductAttributeCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductAttributeCategory 添加产品属性分类表
func (l *AddProductAttributeCategoryLogic) AddProductAttributeCategory(in *pmsclient.AddProductAttributeCategoryReq) (*pmsclient.AddProductAttributeCategoryResp, error) {
	err := query.PmsProductAttributeCategory.WithContext(l.ctx).Create(&model.PmsProductAttributeCategory{
		Name:           in.Name,
		AttributeCount: 0,
		ParamCount:     0,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AddProductAttributeCategoryResp{}, nil
}
