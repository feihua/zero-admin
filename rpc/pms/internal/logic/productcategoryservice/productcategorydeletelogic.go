package productcategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryDeleteLogic 商品类别
/*
Author: LiuFeiHua
Date: 2024/5/8 9:58
*/
type ProductCategoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryDeleteLogic {
	return &ProductCategoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductCategoryDelete 删除商品类别
func (l *ProductCategoryDeleteLogic) ProductCategoryDelete(in *pmsclient.ProductCategoryDeleteReq) (*pmsclient.ProductCategoryDeleteResp, error) {
	q := query.PmsProductCategory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductCategoryDeleteResp{}, nil
}
