package category

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"strconv"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryPcProductCateListLogic
/*
Author: LiuFeiHua
Date: 2024/3/29 下午4:20
*/
type QueryPcProductCateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPcProductCateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPcProductCateListLogic {
	return &QueryPcProductCateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryPcProductCateList 查询所有分类(嵌套children)(pc web端专用)
func (l *QueryPcProductCateListLogic) QueryPcProductCateList() (resp *types.QueryPCProductCateListResp, err error) {
	categoryListResp, err := l.svcCtx.ProductCategoryService.QueryProductCategoryList(l.ctx, &pmsclient.QueryProductCategoryListReq{})

	var list []types.PCCategoryData

	for _, item := range categoryListResp.List {

		var children []types.PCCategoryData
		for _, child := range item.Children {
			children = append(children, types.PCCategoryData{
				Key:   strconv.FormatInt(child.Id, 10),
				Label: child.Name,
			})
		}

		list = append(list, types.PCCategoryData{
			Key:      strconv.FormatInt(item.Id, 10),
			Label:    item.Name,
			Children: children,
		})
	}

	return &types.QueryPCProductCateListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
