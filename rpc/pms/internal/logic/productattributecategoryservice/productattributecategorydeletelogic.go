package productattributecategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeCategoryDeleteLogic 商品属性类别
/*
Author: LiuFeiHua
Date: 2024/5/8 9:53
*/
type ProductAttributeCategoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryDeleteLogic {
	return &ProductAttributeCategoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAttributeCategoryDelete 删除商品属性类别
func (l *ProductAttributeCategoryDeleteLogic) ProductAttributeCategoryDelete(in *pmsclient.ProductAttributeCategoryDeleteReq) (*pmsclient.ProductAttributeCategoryDeleteResp, error) {
	q := query.PmsProductAttributeCategory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeCategoryDeleteResp{}, nil
}
