package productattributecategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryListLogic {
	return &ProductAttributeCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeCategoryListLogic) ProductAttributeCategoryList(in *pmsclient.ProductAttributeCategoryListReq) (*pmsclient.ProductAttributeCategoryListResp, error) {
	q := query.PmsProductAttributeCategory.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.PmsProductAttributeCategory.Name.Like("%" + in.Name + "%"))
	}

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

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

	logc.Infof(l.ctx, "查询商品属性类别列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.ProductAttributeCategoryListResp{
		Total: count,
		List:  list,
	}, nil
}
