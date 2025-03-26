package productattributecategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductAttributeCategoryListLogic 查询产品属性分类列表
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
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

// QueryProductAttributeCategoryList 查询产品属性分类列表
func (l *QueryProductAttributeCategoryListLogic) QueryProductAttributeCategoryList(in *pmsclient.QueryProductAttributeCategoryListReq) (*pmsclient.QueryProductAttributeCategoryListResp, error) {
	productAttributeCategory := query.PmsProductAttributeCategory
	q := productAttributeCategory.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(productAttributeCategory.Name.Like("%" + in.Name + "%"))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询产品属性分类列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询产品属性分类列表失败")
	}

	var list []*pmsclient.ProductAttributeCategoryListData

	for _, item := range result {
		list = append(list, &pmsclient.ProductAttributeCategoryListData{
			Id:             item.ID,             //
			Name:           item.Name,           // 商品属性分类名称
			AttributeCount: item.AttributeCount, // 属性数量
			ParamCount:     item.ParamCount,     // 参数数量

		})
	}

	return &pmsclient.QueryProductAttributeCategoryListResp{
		Total: count,
		List:  list,
	}, nil
}
