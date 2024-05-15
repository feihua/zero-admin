package productattributecategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCategoryWithAttrListLogic
/*
Author: LiuFeiHua
Date: 2024/5/15 14:13
*/
type QueryCategoryWithAttrListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCategoryWithAttrListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCategoryWithAttrListLogic {
	return &QueryCategoryWithAttrListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCategoryWithAttrList 获取所有商品属性分类及其下属性
func (l *QueryCategoryWithAttrListLogic) QueryCategoryWithAttrList(in *pmsclient.ProductAttributeCategoryListReq) (*pmsclient.ProductAttributeCategoryListResp, error) {
	result, err := query.PmsProductAttributeCategory.WithContext(l.ctx).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品属性类别列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductAttributeCategoryListData
	for _, item := range result {
		attribute := query.PmsProductAttribute
		attributes, _ := attribute.WithContext(l.ctx).Where(attribute.ProductAttributeCategoryID.Eq(item.ID)).Find()

		var attrList []*pmsclient.ProductAttributeListData
		for _, r := range attributes {
			attrList = append(attrList, &pmsclient.ProductAttributeListData{
				Id:                         r.ID,
				ProductAttributeCategoryId: r.ProductAttributeCategoryID,
				Name:                       r.Name,
				SelectType:                 r.SelectType,
				InputType:                  r.InputType,
				InputList:                  r.InputList,
				Sort:                       r.Sort,
				FilterType:                 r.FilterType,
				SearchType:                 r.SearchType,
				RelatedStatus:              r.RelatedStatus,
				HandAddStatus:              r.HandAddStatus,
				Type:                       r.Type,
			})
		}
		list = append(list, &pmsclient.ProductAttributeCategoryListData{
			Id:                   item.ID,
			Name:                 item.Name,
			AttributeCount:       item.AttributeCount,
			ParamCount:           item.ParamCount,
			ProductAttributeList: attrList,
		})
	}

	logc.Infof(l.ctx, "查询商品属性类别列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.ProductAttributeCategoryListResp{
		Total: 0,
		List:  list,
	}, nil
}
