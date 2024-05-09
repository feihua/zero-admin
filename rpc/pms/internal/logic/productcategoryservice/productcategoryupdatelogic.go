package productcategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryUpdateLogic 商品类别
/*
Author: LiuFeiHua
Date: 2024/5/8 10:59
*/
type ProductCategoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryUpdateLogic {
	return &ProductCategoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductCategoryUpdate 更新商品类别
func (l *ProductCategoryUpdateLogic) ProductCategoryUpdate(in *pmsclient.ProductCategoryUpdateReq) (*pmsclient.ProductCategoryUpdateResp, error) {
	q := query.PmsProductCategory
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductCategory{
		ID:           in.Id,
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

	return &pmsclient.ProductCategoryUpdateResp{}, nil
}
