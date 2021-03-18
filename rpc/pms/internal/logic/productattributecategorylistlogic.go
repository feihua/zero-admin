package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryListLogic {
	return &ProductAttributeCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeCategoryListLogic) ProductAttributeCategoryList(in *pms.ProductAttributeCategoryListReq) (*pms.ProductAttributeCategoryListResp, error) {
	all, _ := l.svcCtx.PmsProductAttributeCategoryModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsProductAttributeCategoryModel.Count()

	var list []*pms.ProductAttributeCategoryListData
	for _, item := range *all {

		list = append(list, &pms.ProductAttributeCategoryListData{
			Id:             item.Id,
			Name:           item.Name,
			AttributeCount: item.AttributeCount,
			ParamCount:     item.ParamCount,
		})
	}

	fmt.Println(list)
	return &pms.ProductAttributeCategoryListResp{
		Total: count,
		List:  list,
	}, nil
}
