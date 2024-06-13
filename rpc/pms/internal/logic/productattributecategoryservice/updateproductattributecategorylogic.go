package productattributecategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductAttributeCategoryLogic 更新产品属性分类表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:44
*/
type UpdateProductAttributeCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductAttributeCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeCategoryLogic {
	return &UpdateProductAttributeCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductAttributeCategory 更新产品属性分类表
func (l *UpdateProductAttributeCategoryLogic) UpdateProductAttributeCategory(in *pmsclient.UpdateProductAttributeCategoryReq) (*pmsclient.UpdateProductAttributeCategoryResp, error) {
	q := query.PmsProductAttributeCategory
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductAttributeCategory{
		ID:   in.Id,
		Name: in.Name,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.UpdateProductAttributeCategoryResp{}, nil
}
