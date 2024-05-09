package productattributecategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeCategoryAddLogic 商品属性类别
/*
Author: LiuFeiHua
Date: 2024/5/8 10:49
*/
type ProductAttributeCategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryAddLogic {
	return &ProductAttributeCategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAttributeCategoryAdd 添加商品属性类别
func (l *ProductAttributeCategoryAddLogic) ProductAttributeCategoryAdd(in *pmsclient.ProductAttributeCategoryAddReq) (*pmsclient.ProductAttributeCategoryAddResp, error) {
	err := query.PmsProductAttributeCategory.WithContext(l.ctx).Create(&model.PmsProductAttributeCategory{
		Name:           in.Name,
		AttributeCount: 0,
		ParamCount:     0,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeCategoryAddResp{}, nil
}
