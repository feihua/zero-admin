package productcategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductCategoryTreeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCategoryTreeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCategoryTreeListLogic {
	return &QueryProductCategoryTreeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductCategoryTreeList 查询商品分类（tree）
func (l *QueryProductCategoryTreeListLogic) QueryProductCategoryTreeList(in *pmsclient.QueryProductCategoryTreeListReq) (*pmsclient.QueryProductCategoryListTreeResp, error) {
	// 1.查询第一级分类
	categoryList := queryLevel2(l, 0)
	var list []*pmsclient.QueryProductCategoryListTreeData
	for _, item := range categoryList {
		// 2.查询第二级分类
		children := queryLevel2(l, item.Id)
		list = append(list, &pmsclient.QueryProductCategoryListTreeData{
			Id:       item.Id,
			Name:     item.Name,
			ImageUrl: item.ImageUrl,
			Children: children,
		})
	}

	return &pmsclient.QueryProductCategoryListTreeResp{
		List: list,
	}, nil
}

// 查询分类
func queryLevel2(l *QueryProductCategoryTreeListLogic, id int64) []*pmsclient.QueryProductCategoryListTreeData {
	q := query.PmsProductCategory.WithContext(l.ctx)

	q = q.Where(query.PmsProductCategory.IsEnabled.Eq(1))
	if id != 2000 {
		q = q.Where(query.PmsProductCategory.ParentID.Eq(id))
	}

	level2categoryList, _ := q.Find()

	var list []*pmsclient.QueryProductCategoryListTreeData
	for _, category := range level2categoryList {
		list = append(list, &pmsclient.QueryProductCategoryListTreeData{
			Id:       category.ID,
			Name:     category.Name,
			ImageUrl: category.Icon,
		})
	}

	return list
}
