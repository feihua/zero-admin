package category

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCateListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 16:44
*/
type ProductCateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCateListLogic {
	return &ProductCateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductCateList 查询商品分类
func (l *ProductCateListLogic) ProductCateList(req *types.CategoryReq) (resp *types.CategoryResp, err error) {
	categoryListResp, err := l.svcCtx.ProductCategoryService.ProductCategoryList(l.ctx, &pmsclient.ProductCategoryListReq{
		Current:    0,
		PageSize:   100,
		ParentId:   req.ParentId,
		ShowStatus: 1, // 显示状态：0->不显示；1->显示
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logc.Errorf(l.ctx, "参数: %s,查询商品分类列表异常:%s", string(data), err.Error())
		return nil, errors.New("查询商品分类失败")
	}

	list := make([]types.CategoryData, 0)
	for _, item := range categoryListResp.List {
		list = append(list, types.CategoryData{
			Id:       item.Id,
			Name:     item.Name,
			ImageUrl: item.Icon,
		})
	}

	return &types.CategoryResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
