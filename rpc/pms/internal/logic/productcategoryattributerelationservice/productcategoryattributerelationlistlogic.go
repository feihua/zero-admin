package productcategoryattributerelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	q := query.PmsProductCategoryAttributeRelation.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询产品分类和属性关糸列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductCategoryAttributeRelationListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductCategoryAttributeRelationListData{
			Id:                 item.ID,
			ProductCategoryId:  item.ProductCategoryID,
			ProductAttributeId: item.ProductAttributeID,
		})
	}

	logc.Infof(l.ctx, "查询产品分类和属性关糸列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.ProductCategoryAttributeRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
