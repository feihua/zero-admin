package productattributecategoryservicelogic

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

// UpdateProductAttributeCategoryLogic 更新产品属性分类
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
type UpdateProductAttributeCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductAttributeCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeCategoryLogic {
	return &UpdateProductAttributeCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductAttributeCategory 更新产品属性分类
func (l *UpdateProductAttributeCategoryLogic) UpdateProductAttributeCategory(in *pmsclient.UpdateProductAttributeCategoryReq) (*pmsclient.UpdateProductAttributeCategoryResp, error) {
	q := query.PmsProductAttributeCategory.WithContext(l.ctx)

	// 1.根据产品属性分类id查询产品属性分类是否已存在
	_, err := q.Where(query.PmsProductAttributeCategory.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据产品属性分类id：%d,查询产品属性分类失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询产品属性分类失败"))
	}

	item := &model.PmsProductAttributeCategory{
		ID:             in.Id,             //
		Name:           in.Name,           // 商品属性分类名称
		AttributeCount: in.AttributeCount, // 属性数量
		ParamCount:     in.ParamCount,     // 参数数量
	}

	// 2.产品属性分类存在时,则直接更新产品属性分类
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新产品属性分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新产品属性分类失败")
	}

	return &pmsclient.UpdateProductAttributeCategoryResp{}, nil
}
