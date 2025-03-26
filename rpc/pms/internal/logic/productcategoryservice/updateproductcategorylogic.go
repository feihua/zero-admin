package productcategoryservicelogic

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

// UpdateProductCategoryLogic 更新产品分类
/*
Author: LiuFeiHua
Date: 2024/6/12 17:01
*/
type UpdateProductCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductCategoryLogic {
	return &UpdateProductCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductCategory 更新产品分类
func (l *UpdateProductCategoryLogic) UpdateProductCategory(in *pmsclient.UpdateProductCategoryReq) (*pmsclient.UpdateProductCategoryResp, error) {
	// 更新商品分类时要更新商品中的名称
	product := query.PmsProduct
	_, err := product.WithContext(l.ctx).Where(product.ProductCategoryID.Eq(in.Id)).Update(product.ProductCategoryName, in.Name)

	if err != nil {
		logc.Errorf(l.ctx, "更新产品分类失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新产品分类失败")
	}
	// 同时更新筛选属性的信息
	relation := query.PmsProductCategoryAttributeRelation
	_, _ = relation.WithContext(l.ctx).Where(relation.ProductCategoryID.Eq(in.Id)).Delete()
	if len(in.ProductAttributeIdList) > 0 {
		var list []*model.PmsProductCategoryAttributeRelation
		for _, productAttributeId := range in.ProductAttributeIdList {
			list = append(list, &model.PmsProductCategoryAttributeRelation{
				ProductCategoryID:  in.Id,
				ProductAttributeID: productAttributeId,
			})
		}
		_ = relation.WithContext(l.ctx).CreateInBatches(list, len(list))
	}

	q := query.PmsProductCategory
	_, err = q.WithContext(l.ctx).Updates(&model.PmsProductCategory{
		ID:           in.Id,
		ParentID:     in.ParentId,     // 上机分类的编号：0表示一级分类
		Name:         in.Name,         // 商品分类名称
		Level:        in.Level,        // 分类级别：0->1级；1->2级
		ProductCount: in.ProductCount, // 商品数量
		ProductUnit:  in.ProductUnit,  // 商品数量
		NavStatus:    in.NavStatus,    // 是否显示在导航栏：0->不显示；1->显示
		ShowStatus:   in.ShowStatus,   // 显示状态：0->不显示；1->显示
		Sort:         in.Sort,         // 排序
		Icon:         in.Icon,         // 图标
		Keywords:     in.Keywords,     // 关键字
		Description:  in.Description,  // 描述
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新产品分类失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新产品分类失败")
	}

	return &pmsclient.UpdateProductCategoryResp{}, nil
}
