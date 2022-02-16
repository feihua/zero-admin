package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryListLogic {
	return ProductCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryListLogic) ProductCategoryList(req types.ListProductCategoryReq) (*types.ListProductCategoryResp, error) {
	resp, err := l.svcCtx.Pms.ProductCategoryList(l.ctx, &pmsclient.ProductCategoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询商品分类列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询商品分类失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtProductCategoryData

	for _, item := range resp.List {
		list = append(list, &types.ListtProductCategoryData{
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

	return &types.ListProductCategoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询商品分类成功",
	}, nil
}
