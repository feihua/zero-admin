package productattributevalueservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductAttributeValueLogic 更新存储产品参数信息的表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:51
*/
type UpdateProductAttributeValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductAttributeValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeValueLogic {
	return &UpdateProductAttributeValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductAttributeValue 更新存储产品参数信息的表
func (l *UpdateProductAttributeValueLogic) UpdateProductAttributeValue(in *pmsclient.UpdateProductAttributeValueReq) (*pmsclient.UpdateProductAttributeValueResp, error) {
	q := query.PmsProductAttributeValue
	_, err := q.WithContext(l.ctx).Updates(&model.PmsProductAttributeValue{
		ID:                 in.Id,
		ProductID:          in.ProductId,
		ProductAttributeID: in.ProductAttributeId,
		Value:              in.Value,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.UpdateProductAttributeValueResp{}, nil
}
