package productattributevalueservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeValueDeleteLogic 产品参数值
/*
Author: LiuFeiHua
Date: 2024/5/8 9:56
*/
type ProductAttributeValueDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueDeleteLogic {
	return &ProductAttributeValueDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAttributeValueDelete 删除产品参数值
func (l *ProductAttributeValueDeleteLogic) ProductAttributeValueDelete(in *pmsclient.ProductAttributeValueDeleteReq) (*pmsclient.ProductAttributeValueDeleteResp, error) {
	q := query.PmsProductAttributeValue
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeValueDeleteResp{}, nil
}
