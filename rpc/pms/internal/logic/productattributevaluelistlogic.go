package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductAttributeValueListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueListLogic {
	return &ProductAttributeValueListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeValueListLogic) ProductAttributeValueList(in *pms.ProductAttributeValueListReq) (*pms.ProductAttributeValueListResp, error) {
	all, _ := l.svcCtx.PmsProductAttributeValueModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*pms.ProductAttributeValueListData
	for _, item := range *all {

		list = append(list, &pms.ProductAttributeValueListData{
			Id:                 item.Id,
			ProductId:          item.ProductId,
			ProductAttributeId: item.ProductAttributeId,
			Value:              item.Value,
		})
	}

	fmt.Println(list)
	return &pms.ProductAttributeValueListResp{
		Total: 10,
		List:  list,
	}, nil
}
