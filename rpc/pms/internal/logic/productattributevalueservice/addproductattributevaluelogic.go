package productattributevalueservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductAttributeValueLogic 添加商品属性值
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
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

// AddProductAttributeValue 添加商品属性值
func (l *AddProductAttributeValueLogic) AddProductAttributeValue(in *pmsclient.AddProductAttributeValueReq) (*pmsclient.AddProductAttributeValueResp, error) {
	q := query.PmsProductAttributeValue

	item := &model.PmsProductAttributeValue{
		SpuID:       in.SpuId,       // 商品SPU ID
		AttributeID: in.AttributeId, // 属性ID
		Value:       in.Value,       // 属性值
		Status:      in.Status,      // 状态：0->禁用；1->启用
		CreateBy:    in.CreateBy,    // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加商品属性值失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加商品属性值失败")
	}

	return &pmsclient.AddProductAttributeValueResp{}, nil
}
