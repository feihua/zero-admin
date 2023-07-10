package category

import (
	"context"
	"encoding/json"
	"errors"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ProductCateListLogic) ProductCateList(req *types.CategoryReq) (resp *types.CategoryResp, err error) {
	categoryListResp, err := l.svcCtx.ProductCategoryService.ProductCategoryList(l.ctx, &pmsclient.ProductCategoryListReq{
		Current:    0,
		PageSize:   100,
		ParentId:   req.ParentId,
		ShowStatus: 1, // 显示状态：0->不显示；1->显示
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询商品分类列表异常:%s", string(data), err.Error())
		return nil, errors.New("查询商品分类失败")
	}

	var list []types.CategoryData

	for _, item := range categoryListResp.List {
		list = append(list, types.CategoryData{
			Id:           item.Id,
			ParentId:     item.ParentId,
			Name:         item.Name,
			Level:        item.Level,
			ProductCount: item.ProductCount,
			//ProductUnit:  strconv.FormatInt( item.ProductCount,10)+item.ProductUnit,
			ProductUnit: item.ProductUnit,
			NavStatus:   item.NavStatus,
			ShowStatus:  item.ShowStatus,
			Sort:        item.Sort,
			Icon:        item.Icon,
			Keywords:    item.Keywords,
			Description: item.Description,
		})
	}

	return &types.CategoryResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
