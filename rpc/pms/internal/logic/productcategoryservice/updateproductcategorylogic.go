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
	"gorm.io/gorm"
	"time"
)

// UpdateProductCategoryLogic 更新产品分类
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
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
	category := query.PmsProductCategory
	q := category.WithContext(l.ctx)

	// 1.根据产品分类id查询产品分类是否已存在
	detail, err := q.Where(category.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "产品分类不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("产品分类不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询产品分类异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询产品分类异常")
	}

	count, _ := q.WithContext(l.ctx).Where(category.ID.Neq(in.Id), category.Name.Eq(in.Name)).Count()

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("商品分类名称：%s,已存在", in.Name))
	}
	now := time.Now()
	item := &model.PmsProductCategory{
		ID:       in.Id,       //
		ParentID: in.ParentId, // 上级分类的编号：0表示一级分类
		Name:     in.Name,     // 商品分类名称
		Level:    in.Level,    // 分类级别：0->1级；1->2级
		// ProductCount: in.ProductCount, // 商品数量
		ProductUnit: in.ProductUnit,    // 商品单位
		NavStatus:   in.NavStatus,      // 是否显示在导航栏：0->不显示；1->显示
		Sort:        in.Sort,           // 排序
		Icon:        in.Icon,           // 图标
		Keywords:    in.Keywords,       // 关键字
		Description: in.Description,    // 描述
		IsEnabled:   in.IsEnabled,      // 是否启用
		CreateBy:    detail.CreateBy,   // 创建者
		CreateTime:  detail.CreateTime, // 创建时间
		UpdateBy:    &in.UpdateBy,      // 更新者
		UpdateTime:  &now,              // 更新时间
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

	// 2.产品分类存在时,则直接更新产品分类
	err = l.svcCtx.DB.Model(&model.PmsProductCategory{}).WithContext(l.ctx).Where(category.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新产品分类失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新产品分类失败")
	}

	return &pmsclient.UpdateProductCategoryResp{}, nil
}
