package productcategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryListLogic 商品类别
/*
Author: LiuFeiHua
Date: 2023/11/29 16:45
*/
type ProductCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryListLogic {
	return &ProductCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductCategoryList 查询商品分类列表
func (l *ProductCategoryListLogic) ProductCategoryList(in *pmsclient.ProductCategoryListReq) (*pmsclient.ProductCategoryListResp, error) {
	q := query.PmsProductCategory.WithContext(l.ctx)

	if in.ParentId != 2000 {
		q = q.Where(query.PmsProductCategory.ParentID.Eq(in.ParentId))
	}
	if in.ShowStatus != 2 {
		q = q.Where(query.PmsProductCategory.ShowStatus.Eq(in.ShowStatus))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品类别列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductCategoryListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductCategoryListData{
			Id:           item.ID,
			ParentId:     item.ParentID,
			Name:         item.Name,
			Level:        item.Level,
			ProductCount: item.ProductCount,
			ProductUnit:  item.ProductUnit,
			NavStatus:    item.NavStatus,
			ShowStatus:   item.ShowStatus,
			Sort:         item.Sort,
			Icon:         item.Icon,
			Keywords:     item.Keywords,
			Description:  *item.Description,
		})
	}

	logc.Infof(l.ctx, "查询商品类别列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.ProductCategoryListResp{
		Total: count,
		List:  list,
	}, nil
}
