package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryListLogic {
	return &ProductCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryListLogic) ProductCategoryList(in *pms.ProductCategoryListReq) (*pms.ProductCategoryListResp, error) {
	all, _ := l.svcCtx.PmsProductCategoryModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*pms.ProductCategoryListData
	for _, item := range *all {

		list = append(list, &pms.ProductCategoryListData{
			Id:           item.Id,
			ParentId:     item.ParentId,
			Name:         item.Name,
			Level:        item.Level,
			ProductCount: item.ProductCount,
			ProductUnit:  item.ProductUnit,
			NavStatus:    item.NavStatus,
			ShowStatus:   item.ShowStatus,
			Sort:         item.Sort,
			Icon:         item.Icon,
			Keywords:     item.Keywords,
			Description:  item.Description,
		})
	}

	fmt.Println(list)
	return &pms.ProductCategoryListResp{
		Total: 10,
		List:  list,
	}, nil
}
