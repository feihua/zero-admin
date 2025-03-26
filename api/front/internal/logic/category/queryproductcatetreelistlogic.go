package category

import (
	"context"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strconv"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCateTreeListLogic 以树形结构获取所有商品分类
/*
Author: LiuFeiHua
Date: 2024/5/16 14:50
*/
type QueryProductCateTreeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductCateTreeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCateTreeListLogic {
	return &QueryProductCateTreeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductCateTreeList 以树形结构获取所有商品分类
func (l *QueryProductCateTreeListLogic) QueryProductCateTreeList() (resp *types.QueryProductCateListResp, err error) {
	categoryListResp, err := l.svcCtx.ProductCategoryService.QueryProductCategoryTreeList(l.ctx, &pmsclient.QueryProductCategoryTreeListReq{})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品分类列表失败,异常：%s", err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []types.ProductCateListData

	for _, item := range categoryListResp.List {

		var children []types.ProductCateListData
		for _, child := range item.Children {
			children = append(children, types.ProductCateListData{
				Id:       child.Id,
				Key:      strconv.FormatInt(child.Id, 10),
				Label:    child.Name,
				Name:     child.Name,
				ImageUrl: child.ImageUrl,
			})
		}

		list = append(list, types.ProductCateListData{
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
