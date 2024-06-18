package attributecategory

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
				Id:                         i.Id,
				ProductAttributeCategoryId: i.ProductAttributeCategoryId,
				Name:                       i.Name,
				SelectType:                 i.SelectType,
				InputType:                  i.InputType,
				InputList:                  i.InputList,
				Sort:                       i.Sort,
				FilterType:                 i.FilterType,
				SearchType:                 i.SearchType,
				RelatedStatus:              i.RelatedStatus,
				HandAddStatus:              i.HandAddStatus,
				Type:                       i.Type,
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
