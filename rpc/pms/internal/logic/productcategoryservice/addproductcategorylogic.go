package productcategoryservicelogic

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

// AddProductCategoryLogic 添加产品分类
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
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
	q := query.PmsProductCategory
	count, _ := q.WithContext(l.ctx).Where(q.Name.Eq(in.Name)).Count()

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("商品分类名称：%s,已存在", in.Name))
	}
	item := &model.PmsProductCategory{
		ParentID:     in.ParentId,    // 上级分类的编号：0表示一级分类
		Name:         in.Name,        // 商品分类名称
		Level:        in.Level,       // 分类级别：0->1级；1->2级
		ProductCount: 0,              // 商品数量
		ProductUnit:  in.ProductUnit, // 商品单位
		NavStatus:    in.NavStatus,   // 是否显示在导航栏：0->不显示；1->显示
		Sort:         in.Sort,        // 排序
		Icon:         in.Icon,        // 图标
		Keywords:     in.Keywords,    // 关键字
		Description:  in.Description, // 描述
		IsEnabled:    in.IsEnabled,   // 是否启用
		CreateBy:     in.CreateBy,    // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加产品分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加产品分类失败")
	}
	if len(in.ProductAttributeIdList) > 0 {
		var list []*model.PmsProductCategoryAttributeRelation
		for _, productAttributeId := range in.ProductAttributeIdList {
			list = append(list, &model.PmsProductCategoryAttributeRelation{
				ProductCategoryID:  item.ID,
				ProductAttributeID: productAttributeId,
			})
		}

		_ = query.PmsProductCategoryAttributeRelation.WithContext(l.ctx).CreateInBatches(list, len(list))
	}
	return &pmsclient.AddProductCategoryResp{}, nil
}
