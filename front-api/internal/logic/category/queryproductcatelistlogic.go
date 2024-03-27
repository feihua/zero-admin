package category

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"strconv"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCateListLogic
/*
Author: LiuFeiHua
Date: 2024/3/7 下午2:24
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

// QueryProductCateList 查询所有分类
func (l *QueryProductCateListLogic) QueryProductCateList() (resp *types.QueryProductCateListResp, err error) {
	categoryListResp, err := l.svcCtx.ProductCategoryService.QueryProductCategoryList(l.ctx, &pmsclient.QueryProductCategoryListReq{})

	var list []types.CategoryData

	for _, item := range categoryListResp.List {

		var children []types.CategoryData
		for _, child := range item.Children {
			children = append(children, types.CategoryData{
				Id:       child.Id,
				Key:      strconv.FormatInt(child.Id, 10),
				Label:    child.Name,
				Name:     child.Name,
				ImageUrl: child.ImageUrl,
			})
		}

		list = append(list, types.CategoryData{
			Id:       item.Id,
			Name:     item.Name,
			Key:      strconv.FormatInt(item.Id, 10),
			Label:    item.Name,
			ImageUrl: item.ImageUrl,
			Children: children,
		})
	}

	return &types.QueryProductCateListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
