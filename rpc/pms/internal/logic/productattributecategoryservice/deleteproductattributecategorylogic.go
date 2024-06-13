package productattributecategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductAttributeCategoryLogic 删除产品属性分类表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:42
*/
type DeleteProductAttributeCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductAttributeCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductAttributeCategoryLogic {
	return &DeleteProductAttributeCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductAttributeCategory 删除产品属性分类表
func (l *DeleteProductAttributeCategoryLogic) DeleteProductAttributeCategory(in *pmsclient.DeleteProductAttributeCategoryReq) (*pmsclient.DeleteProductAttributeCategoryResp, error) {
	q := query.PmsProductAttributeCategory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.DeleteProductAttributeCategoryResp{}, nil
}
