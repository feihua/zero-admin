package productattributevalueservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeValueAddLogic 属性值
/*
Author: LiuFeiHua
Date: 2024/5/8 10:53
*/
type ProductAttributeValueAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueAddLogic {
	return &ProductAttributeValueAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductAttributeValueAdd 添加属性值
func (l *ProductAttributeValueAddLogic) ProductAttributeValueAdd(in *pmsclient.ProductAttributeValueAddReq) (*pmsclient.ProductAttributeValueAddResp, error) {
	err := query.PmsProductAttributeValue.WithContext(l.ctx).Create(&model.PmsProductAttributeValue{
		ProductID:          in.ProductId,
		ProductAttributeID: in.ProductAttributeId,
		Value:              &in.Value,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductAttributeValueAddResp{}, nil
}
