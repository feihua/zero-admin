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

// AddProductCategoryLogic 添加产品分类
/*
Author: LiuFeiHua
Date: 2024/6/12 16:53
*/
type AddProductCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductCategoryLogic {
	return &AddProductCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductCategory 添加产品分类
func (l *AddProductCategoryLogic) AddProductCategory(in *pmsclient.AddProductCategoryReq) (*pmsclient.AddProductCategoryResp, error) {
	category := &model.PmsProductCategory{
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
	}
	err := query.PmsProductCategory.WithContext(l.ctx).Create(category)

	if err != nil {
		logc.Errorf(l.ctx, "添加产品分类失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加产品分类失败")
	}

	if len(in.ProductAttributeIdList) > 0 {
		var list []*model.PmsProductCategoryAttributeRelation
		for _, productAttributeId := range in.ProductAttributeIdList {
			list = append(list, &model.PmsProductCategoryAttributeRelation{
				ProductCategoryID:  category.ID,
				ProductAttributeID: productAttributeId,
			})
		}

		_ = query.PmsProductCategoryAttributeRelation.WithContext(l.ctx).CreateInBatches(list, len(list))
	}

	return &pmsclient.AddProductCategoryResp{}, nil
}
