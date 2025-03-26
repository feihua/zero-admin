package productattributevalueservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductAttributeValueLogic 添加存储产品参数信息
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

// AddProductAttributeValue 添加存储产品参数信息
func (l *AddProductAttributeValueLogic) AddProductAttributeValue(in *pmsclient.AddProductAttributeValueReq) (*pmsclient.AddProductAttributeValueResp, error) {
	err := query.PmsProductAttributeValue.WithContext(l.ctx).Create(&model.PmsProductAttributeValue{
		ProductID:          in.ProductId,          // 商品id
		ProductAttributeID: in.ProductAttributeId, // 商品属性id
		Value:              in.Value,              // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加存储产品参数信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加存储产品参数信息失败")
	}

	return &pmsclient.AddProductAttributeValueResp{}, nil
}
