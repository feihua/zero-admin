package productattributecategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductAttributeCategoryListLogic 查询产品属性分类表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:43
*/
type QueryProductAttributeCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductAttributeCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeCategoryListLogic {
	return &QueryProductAttributeCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductAttributeCategoryList 查询产品属性分类表列表
func (l *QueryProductAttributeCategoryListLogic) QueryProductAttributeCategoryList(in *pmsclient.QueryProductAttributeCategoryListReq) (*pmsclient.QueryProductAttributeCategoryListResp, error) {
	q := query.PmsProductAttributeCategory.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.PmsProductAttributeCategory.Name.Like("%" + in.Name + "%"))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品属性类别列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductAttributeCategoryListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductAttributeCategoryListData{
			Id:             item.ID,
			Name:           item.Name,
			AttributeCount: item.AttributeCount,
			ParamCount:     item.ParamCount,
		})
	}

	return &pmsclient.QueryProductAttributeCategoryListResp{
		Total: count,
		List:  list,
	}, nil

}
