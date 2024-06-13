package productattributeservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductAttributeListLogic 查询商品属性参数表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:48
*/
type QueryProductAttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeListLogic {
	return &QueryProductAttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductAttributeList 查询商品属性参数表列表
func (l *QueryProductAttributeListLogic) QueryProductAttributeList(in *pmsclient.QueryProductAttributeListReq) (*pmsclient.QueryProductAttributeListResp, error) {
	q := query.PmsProductAttribute.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.PmsProductAttribute.Name.Like("%" + in.Name + "%"))
	}

	if in.Type != 2 {
		q = q.Where(query.PmsProductAttribute.Type.Eq(in.Type))
	}

	if in.ProductAttributeCategoryId != 0 {
		q = q.Where(query.PmsProductAttribute.ProductAttributeCategoryID.Eq(in.ProductAttributeCategoryId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品属性列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductAttributeListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductAttributeListData{
			Id:                         item.ID,
			ProductAttributeCategoryId: item.ProductAttributeCategoryID,
			Name:                       item.Name,
			SelectType:                 item.SelectType,
			InputType:                  item.InputType,
			InputList:                  item.InputList,
			Sort:                       item.Sort,
			FilterType:                 item.FilterType,
			SearchType:                 item.SearchType,
			RelatedStatus:              item.RelatedStatus,
			HandAddStatus:              item.HandAddStatus,
			Type:                       item.Type,
		})
	}

	return &pmsclient.QueryProductAttributeListResp{
		Total: count,
		List:  list,
	}, nil

}
