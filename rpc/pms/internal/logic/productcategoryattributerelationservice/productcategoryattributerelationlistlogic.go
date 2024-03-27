package productcategoryattributerelationservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryAttributeRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationListLogic {
	return &ProductCategoryAttributeRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryAttributeRelationListLogic) ProductCategoryAttributeRelationList(in *pmsclient.ProductCategoryAttributeRelationListReq) (*pmsclient.ProductCategoryAttributeRelationListResp, error) {
	all, err := l.svcCtx.PmsProductCategoryAttributeRelationModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsProductCategoryAttributeRelationModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询产品分类和属性关糸列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductCategoryAttributeRelationListData
	for _, item := range *all {

		list = append(list, &pmsclient.ProductCategoryAttributeRelationListData{
			Id:                 item.Id,
			ProductCategoryId:  item.ProductCategoryId,
			ProductAttributeId: item.ProductAttributeId,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询产品分类和属性关糸列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pmsclient.ProductCategoryAttributeRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
