package productattributecategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeCategoryUpdateLogic 商品属性类别
/*
Author: LiuFeiHua
Date: 2024/5/8 10:50
*/
type ProductAttributeCategoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryUpdateLogic {
	return &ProductAttributeCategoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAttributeCategoryUpdate 更新商品属性类别
func (l *ProductAttributeCategoryUpdateLogic) ProductAttributeCategoryUpdate(in *pmsclient.ProductAttributeCategoryUpdateReq) (*pmsclient.ProductAttributeCategoryUpdateResp, error) {
	q := query.PmsProductAttributeCategory
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductAttributeCategory{
		ID:   in.Id,
		Name: in.Name,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeCategoryUpdateResp{}, nil
}
