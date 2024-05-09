package productattributeservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeDeleteLogic 商品属性
/*
Author: LiuFeiHua
Date: 2024/5/8 9:54
*/
type ProductAttributeDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeDeleteLogic {
	return &ProductAttributeDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAttributeDelete 删除商品属性
func (l *ProductAttributeDeleteLogic) ProductAttributeDelete(in *pmsclient.ProductAttributeDeleteReq) (*pmsclient.ProductAttributeDeleteResp, error) {
	q := query.PmsProductAttribute
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeDeleteResp{}, nil
}
