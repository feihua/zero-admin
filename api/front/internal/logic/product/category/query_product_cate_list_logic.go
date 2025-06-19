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

// QueryProductCateListLogic 获取所有商品分类
/*
Author: LiuFeiHua
Date: 2025/6/19 10:31
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

// QueryProductCateList 获取所有商品分类
func (l *QueryProductCateListLogic) QueryProductCateList() (resp *types.QueryProductCateListResp, err error) {
	categoryListResp, err := l.svcCtx.ProductCategoryService.QueryProductCategoryTreeList(l.ctx, &pmsclient.QueryProductCategoryTreeListReq{})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品分类列表失败,异常：%s", err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []types.ProductCateData

	for _, item := range categoryListResp.List {

		var children []types.ProductCateData
		for _, child := range item.Children {
			children = append(children, types.ProductCateData{
				Id:       child.Id,
				Key:      strconv.FormatInt(child.Id, 10),
				Label:    child.Name,
				Name:     child.Name,
				ImageUrl: child.ImageUrl,
			})
		}

		list = append(list, types.ProductCateData{
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
