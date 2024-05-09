package productcategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryAddLogic 商品类别
/*
Author: LiuFeiHua
Date: 2024/5/8 10:59
*/
type ProductCategoryAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAddLogic {
	return &ProductCategoryAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductCategoryAdd 添加商品类别
func (l *ProductCategoryAddLogic) ProductCategoryAdd(in *pmsclient.ProductCategoryAddReq) (*pmsclient.ProductCategoryAddResp, error) {
	err := query.PmsProductCategory.WithContext(l.ctx).Create(&model.PmsProductCategory{
		ParentID:     in.ParentId,
		Name:         in.Name,
		Level:        in.Level,
		ProductCount: in.ProductCount,
		ProductUnit:  in.ProductUnit,
		NavStatus:    in.NavStatus,
		ShowStatus:   in.ShowStatus,
		Sort:         in.Sort,
		Icon:         in.Icon,
		Keywords:     in.Keywords,
		Description:  &in.Description,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductCategoryAddResp{}, nil
}
