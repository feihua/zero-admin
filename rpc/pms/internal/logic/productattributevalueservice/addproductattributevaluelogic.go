package productattributevalueservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductAttributeValueLogic 添加存储产品参数信息的表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:49
*/
type AddProductAttributeValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductAttributeValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductAttributeValueLogic {
	return &AddProductAttributeValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductAttributeValue 添加存储产品参数信息的表
func (l *AddProductAttributeValueLogic) AddProductAttributeValue(in *pmsclient.AddProductAttributeValueReq) (*pmsclient.AddProductAttributeValueResp, error) {
	err := query.PmsProductAttributeValue.WithContext(l.ctx).Create(&model.PmsProductAttributeValue{
		ProductID:          in.ProductId,
		ProductAttributeID: in.ProductAttributeId,
		Value:              in.Value,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AddProductAttributeValueResp{}, nil
}
