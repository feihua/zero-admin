package productattributevalueservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeValueUpdateLogic 属性值
/*
Author: LiuFeiHua
Date: 2024/5/8 10:54
*/
type ProductAttributeValueUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueUpdateLogic {
	return &ProductAttributeValueUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAttributeValueUpdate 更新属性值
func (l *ProductAttributeValueUpdateLogic) ProductAttributeValueUpdate(in *pmsclient.ProductAttributeValueUpdateReq) (*pmsclient.ProductAttributeValueUpdateResp, error) {
	q := query.PmsProductAttributeValue
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductAttributeValue{
		ID:                 in.Id,
		ProductID:          in.ProductId,
		ProductAttributeID: in.ProductAttributeId,
		Value:              &in.Value,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeValueUpdateResp{}, nil
}
