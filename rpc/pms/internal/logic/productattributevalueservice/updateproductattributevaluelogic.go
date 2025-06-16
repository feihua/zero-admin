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
	"gorm.io/gorm"
	"time"
)

// UpdateProductAttributeValueLogic 更新商品属性值
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
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

// UpdateProductAttributeValue 更新商品属性值
func (l *UpdateProductAttributeValueLogic) UpdateProductAttributeValue(in *pmsclient.UpdateProductAttributeValueReq) (*pmsclient.UpdateProductAttributeValueResp, error) {
	value := query.PmsProductAttributeValue
	q := value.WithContext(l.ctx)

	// 1.根据商品属性值id查询商品属性值是否已存在
	detail, err := q.Where(value.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品属性值不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品属性值不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品属性值异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品属性值异常")
	}

	now := time.Now()
	item := &model.PmsProductAttributeValue{
		ID:          in.Id,             // 主键id
		SpuID:       in.SpuId,          // 商品SPU ID
		AttributeID: in.AttributeId,    // 属性ID
		Value:       in.Value,          // 属性值
		Status:      in.Status,         // 状态：0->禁用；1->启用
		CreateBy:    detail.CreateBy,   // 创建人ID
		CreateTime:  detail.CreateTime, // 创建时间
		UpdateBy:    &in.UpdateBy,      // 更新人ID
		UpdateTime:  &now,              // 更新时间
	}

	// 2.商品属性值存在时,则直接更新商品属性值
	err = l.svcCtx.DB.Model(&model.PmsProductAttributeValue{}).WithContext(l.ctx).Where(value.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新商品属性值失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新商品属性值失败")
	}

	return &pmsclient.UpdateProductAttributeValueResp{}, nil
}
