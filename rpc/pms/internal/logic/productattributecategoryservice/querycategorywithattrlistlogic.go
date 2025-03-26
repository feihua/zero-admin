package productattributecategoryservicelogic

import (
	"context"
	"errors"
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
func (l *QueryCategoryWithAttrListLogic) QueryCategoryWithAttrList(in *pmsclient.QueryProductAttributeCategoryListReq) (*pmsclient.QueryProductAttributeCategoryListResp, error) {
	result, err := query.PmsProductAttributeCategory.WithContext(l.ctx).Find()

	if err != nil {
		logc.Errorf(l.ctx, "获取所有商品属性分类及其下属性失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("获取所有商品属性分类及其下属性失败")
	}

	var list []*pmsclient.ProductAttributeCategoryListData
	for _, item1 := range result {
		attribute := query.PmsProductAttribute
		attributes, _ := attribute.WithContext(l.ctx).Where(attribute.ProductAttributeCategoryID.Eq(item1.ID)).Find()

		var attrList []*pmsclient.ProductAttributeData
		for _, item := range attributes {
			attrList = append(attrList, &pmsclient.ProductAttributeData{
				Id:                         item.ID,                         //
				ProductAttributeCategoryId: item.ProductAttributeCategoryID, // 商品属性分类id
				Name:                       item.Name,                       // 商品属性分类id
				SelectType:                 item.SelectType,                 // 属性选择类型：0->唯一；1->单选；2->多选
				InputType:                  item.InputType,                  // 属性录入方式：0->手工录入；1->从列表中选取
				InputList:                  item.InputList,                  // 可选值列表，以逗号隔开
				Sort:                       item.Sort,                       // 排序字段：最高的可以单独上传图片
				FilterType:                 item.FilterType,                 // 分类筛选样式：1->普通；1->颜色
				SearchType:                 item.SearchType,                 // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
				RelatedStatus:              item.RelatedStatus,              // 相同属性产品是否关联；0->不关联；1->关联
				HandAddStatus:              item.HandAddStatus,              // 是否支持手动新增；0->不支持；1->支持
				Type:                       item.Type,                       // 属性的类型；0->规格；1->参数
			})
		}
		list = append(list, &pmsclient.ProductAttributeCategoryListData{
			Id:                   item1.ID,
			Name:                 item1.Name,
			AttributeCount:       item1.AttributeCount,
			ParamCount:           item1.ParamCount,
			ProductAttributeList: attrList,
		})
	}

	return &pmsclient.QueryProductAttributeCategoryListResp{
		Total: 0,
		List:  list,
	}, nil
}
