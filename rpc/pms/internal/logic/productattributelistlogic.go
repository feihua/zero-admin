package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeListLogic {
	return &ProductAttributeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeListLogic) ProductAttributeList(in *pms.ProductAttributeListReq) (*pms.ProductAttributeListResp, error) {
	all, _ := l.svcCtx.PmsProductAttributeModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*pms.ProductAttributeListData
	for _, item := range *all {

		list = append(list, &pms.ProductAttributeListData{
			Id:                         item.Id,
			ProductAttributeCategoryId: item.ProductAttributeCategoryId,
			Name:                       item.Name,
			SelectType:                 item.SelectType,
			InputType:                  item.InputType,
			InputList:                  item.InputList,
			Sort:                       item.Sort,
			FilterType:                 item.FilterType,
			SearchType:                 item.SearchType,
			RelatedStatus:              item.RelatedStatus,
			HandAddStatus:              item.HandAddStatus,
			Type:                       item.Type,
		})
	}

	fmt.Println(list)
	return &pms.ProductAttributeListResp{
		Total: 10,
		List:  list,
	}, nil
}
