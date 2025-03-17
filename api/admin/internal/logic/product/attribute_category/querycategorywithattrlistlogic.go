package attribute_category

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCategoryWithAttrListLogic 获取所有商品属性分类及其下属性
/*
Author: LiuFeiHua
Date: 2024/5/15 14:09
*/
type QueryCategoryWithAttrListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCategoryWithAttrListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCategoryWithAttrListLogic {
	return &QueryCategoryWithAttrListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryCategoryWithAttrList 获取所有商品属性分类及其下属性
func (l *QueryCategoryWithAttrListLogic) QueryCategoryWithAttrList() (resp *types.ListProductAttributeCategoryResp, err error) {
	attributeList, er := l.svcCtx.ProductAttributeCategoryService.QueryCategoryWithAttrList(l.ctx, &pmsclient.QueryProductAttributeCategoryListReq{})

	if er != nil {
		logc.Errorf(l.ctx, "获取所有商品属性分类及其下属性列表异常:%s", er.Error())
		return nil, errorx.NewDefaultError("获取所有商品属性分类及其下属性失败")
	}

	var list []*types.ListProductAttributeCategoryData

	for _, item := range attributeList.List {
		var attr []types.ListProductAttribute
		for _, i := range item.ProductAttributeList {
			attr = append(attr, types.ListProductAttribute{
				Id:                         i.Id,                         //
				ProductAttributeCategoryId: i.ProductAttributeCategoryId, // 商品属性分类id
				Name:                       i.Name,                       // 商品属性分类id
				SelectType:                 i.SelectType,                 // 属性选择类型：0->唯一；1->单选；2->多选
				InputType:                  i.InputType,                  // 属性录入方式：0->手工录入；1->从列表中选取
				InputList:                  i.InputList,                  // 可选值列表，以逗号隔开
				Sort:                       i.Sort,                       // 排序字段：最高的可以单独上传图片
				FilterType:                 i.FilterType,                 // 分类筛选样式：1->普通；1->颜色
				SearchType:                 i.SearchType,                 // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
				RelatedStatus:              i.RelatedStatus,              // 相同属性产品是否关联；0->不关联；1->关联
				HandAddStatus:              i.HandAddStatus,              // 是否支持手动新增；0->不支持；1->支持
				Type:                       i.Type,                       // 属性的类型；0->规格；1->参数
			})
		}
		list = append(list, &types.ListProductAttributeCategoryData{
			Id:                     item.Id,
			Name:                   item.Name,
			AttributecategoryCount: item.AttributeCount,
			ParamCount:             item.ParamCount,
			ProductAttributeList:   attr,
		})
	}

	return &types.ListProductAttributeCategoryResp{
		Current:  0,
		Data:     list,
		PageSize: 0,
		Success:  true,
		Total:    attributeList.Total,
		Code:     "000000",
		Message:  "获取所有商品属性分类及其下属性成功",
	}, nil
}
