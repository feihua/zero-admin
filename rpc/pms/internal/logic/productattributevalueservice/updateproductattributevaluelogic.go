package productattributevalueservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductAttributeValueLogic 更新存储产品参数信息的
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
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

// UpdateProductAttributeValue 更新存储产品参数信息的
func (l *UpdateProductAttributeValueLogic) UpdateProductAttributeValue(in *pmsclient.UpdateProductAttributeValueReq) (*pmsclient.UpdateProductAttributeValueResp, error) {
	q := query.PmsProductAttributeValue.WithContext(l.ctx)

	// 1.根据存储产品参数信息的id查询存储产品参数信息的是否已存在
	_, err := q.Where(query.PmsProductAttributeValue.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据存储产品参数信息的id：%d,查询存储产品参数信息的失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询存储产品参数信息的失败"))
	}

	item := &model.PmsProductAttributeValue{
		ID:                 in.Id,                 //
		ProductID:          in.ProductId,          // 商品id
		ProductAttributeID: in.ProductAttributeId, // 商品属性id
		Value:              in.Value,              // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
	}

	// 2.存储产品参数信息的存在时,则直接更新存储产品参数信息的
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新存储产品参数信息的失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新存储产品参数信息的失败")
	}

	logc.Infof(l.ctx, "更新存储产品参数信息的成功,参数：%+v", in)
	return &pmsclient.UpdateProductAttributeValueResp{}, nil
}
