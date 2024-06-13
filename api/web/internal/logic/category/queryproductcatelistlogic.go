package category

import (
	"context"
	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCateListLogic
/*
Author: LiuFeiHua
Date: 2024/4/7 17:51
*/
type QueryProductCateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductCateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCateListLogic {
	return &QueryProductCateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductCateList 查询商品分类
func (l *QueryProductCateListLogic) QueryProductCateList() (*types.QueryProductCateListResp, error) {
	categoryListResp, _ := l.svcCtx.ProductCategoryService.QueryProductCategoryTreeList(l.ctx, &pmsclient.QueryProductCategoryTreeListReq{})

	var list []types.CategoryData

	for _, item := range categoryListResp.List {

		var children []types.CategoryData
		for _, child := range item.Children {
			children = append(children, types.CategoryData{
				Id:   child.Id,
				Name: child.Name,
			})
		}

		list = append(list, types.CategoryData{
			Id:       item.Id,
			Name:     item.Name,
			Children: children,
		})
	}

	return &types.QueryProductCateListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
