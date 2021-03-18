package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCategoryAttributeRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationListLogic {
	return &ProductCategoryAttributeRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryAttributeRelationListLogic) ProductCategoryAttributeRelationList(in *pms.ProductCategoryAttributeRelationListReq) (*pms.ProductCategoryAttributeRelationListResp, error) {
	all, _ := l.svcCtx.PmsProductCategoryAttributeRelationModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsProductCategoryAttributeRelationModel.Count()

	var list []*pms.ProductCategoryAttributeRelationListData
	for _, item := range *all {

		list = append(list, &pms.ProductCategoryAttributeRelationListData{
			Id:                 item.Id,
			ProductCategoryId:  item.ProductCategoryId,
			ProductAttributeId: item.ProductAttributeId,
		})
	}

	fmt.Println(list)
	return &pms.ProductCategoryAttributeRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
