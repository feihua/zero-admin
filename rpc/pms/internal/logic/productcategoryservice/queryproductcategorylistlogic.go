package productcategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCategoryListLogic 查询商品类别列表逻辑处理
/*
Author: LiuFeiHua
Date: 2024/3/7 下午2:58
*/
type QueryProductCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCategoryListLogic {
	return &QueryProductCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductCategoryList 查询商品类别列表
func (l *QueryProductCategoryListLogic) QueryProductCategoryList(in *pmsclient.QueryProductCategoryListReq) (*pmsclient.QueryProductCategoryListResp, error) {
	//1.查询第一级分类
	categoryList := queryLevel2(l, 0)
	var list []*pmsclient.QueryProductCategoryListData
	for _, item := range categoryList {
		//2.查询第二级分类
		children := queryLevel2(l, item.Id)
		list = append(list, &pmsclient.QueryProductCategoryListData{
			Id:       item.Id,
			Name:     item.Name,
			ImageUrl: item.ImageUrl,
			Children: children,
		})
	}

	return &pmsclient.QueryProductCategoryListResp{
		List: list,
	}, nil
}

// 查询分类
func queryLevel2(l *QueryProductCategoryListLogic, id int64) []*pmsclient.QueryProductCategoryListData {
	q := query.PmsProductCategory.WithContext(l.ctx)

	q = q.Where(query.PmsProductCategory.ShowStatus.Eq(1))
	if id != 2000 {
		q = q.Where(query.PmsProductCategory.ParentID.Eq(id))
	}

	level2categoryList, _ := q.Find()

	var list []*pmsclient.QueryProductCategoryListData
	for _, category := range level2categoryList {
		list = append(list, &pmsclient.QueryProductCategoryListData{
			Id:       category.ID,
			Name:     category.Name,
			ImageUrl: category.Icon,
		})
	}

	return list
}
