package productattributeservicelogic

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

// AddProductAttributeLogic 添加商品属性参数
/*
Author: LiuFeiHua
Date: 2024/6/12 16:47
*/
type AddProductAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductAttributeLogic {
	return &AddProductAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductAttribute 添加商品属性参数
func (l *AddProductAttributeLogic) AddProductAttribute(in *pmsclient.AddProductAttributeReq) (*pmsclient.AddProductAttributeResp, error) {
	err := query.PmsProductAttribute.WithContext(l.ctx).Create(&model.PmsProductAttribute{
		ProductAttributeCategoryID: in.ProductAttributeCategoryId, // 商品属性分类id
		Name:                       in.Name,                       // 商品属性分类id
		SelectType:                 in.SelectType,                 // 属性选择类型：0->唯一；1->单选；2->多选
		InputType:                  in.InputType,                  // 属性录入方式：0->手工录入；1->从列表中选取
		InputList:                  in.InputList,                  // 可选值列表，以逗号隔开
		Sort:                       in.Sort,                       // 排序字段：最高的可以单独上传图片
		FilterType:                 in.FilterType,                 // 分类筛选样式：1->普通；1->颜色
		SearchType:                 in.SearchType,                 // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
		RelatedStatus:              in.RelatedStatus,              // 相同属性产品是否关联；0->不关联；1->关联
		HandAddStatus:              in.HandAddStatus,              // 是否支持手动新增；0->不支持；1->支持
		Type:                       in.Type,                       // 属性的类型；0->规格；1->参数
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品属性参数失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加商品属性参数失败")
	}

	// 新增商品属性以后需要更新商品属性分类数量
	q := query.PmsProductAttributeCategory
	categoryDo := q.WithContext(l.ctx).Where(q.ID.Eq(in.ProductAttributeCategoryId))
	category, err := categoryDo.First()
	if err != nil {
		logc.Errorf(l.ctx, "添加商品属性参数失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加商品属性参数失败")
	}

	if in.Type == 0 {
		_, _ = categoryDo.Update(q.AttributeCount, category.AttributeCount+1)
	}

	if in.Type == 1 {
		_, _ = categoryDo.Update(q.ParamCount, category.ParamCount+1)
	}

	return &pmsclient.AddProductAttributeResp{}, nil
}
