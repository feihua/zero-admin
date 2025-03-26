package productcategoryattributerelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCategoryAttributeRelationListLogic 查询产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）列表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:52
*/
type QueryProductCategoryAttributeRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCategoryAttributeRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCategoryAttributeRelationListLogic {
	return &QueryProductCategoryAttributeRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductCategoryAttributeRelationList 查询产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）列表
func (l *QueryProductCategoryAttributeRelationListLogic) QueryProductCategoryAttributeRelationList(in *pmsclient.QueryProductCategoryAttributeRelationListReq) (*pmsclient.QueryProductCategoryAttributeRelationListResp, error) {
	q := query.PmsProductCategoryAttributeRelation.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询产品的分类和属性的关系失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询产品的分类和属性的关系失败")
	}

	var list []*pmsclient.ProductCategoryAttributeRelationListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductCategoryAttributeRelationListData{
			Id:                 item.ID,                 //
			ProductCategoryId:  item.ProductCategoryID,  // 商品分类id
			ProductAttributeId: item.ProductAttributeID, // 商品属性id
		})
	}

	return &pmsclient.QueryProductCategoryAttributeRelationListResp{
		Total: count,
		List:  list,
	}, nil

}
